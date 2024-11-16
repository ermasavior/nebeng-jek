package main

import (
	"os/signal"
	"syscall"
	"time"

	"context"
	ridesHandler "nebeng-jek/internal/rides/handler"
	"nebeng-jek/pkg/configs"
	db "nebeng-jek/pkg/db/postgres"
	pkgHttp "nebeng-jek/pkg/http"
	"nebeng-jek/pkg/http_client"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/messaging/nats"
	pkgOtel "nebeng-jek/pkg/otel"
	"os"
)

func main() {
	ctx := context.Background()

	envFilePath := os.Getenv("ENV_PATH")
	cfg := configs.NewConfig(envFilePath)

	otel := pkgOtel.NewOpenTelemetry(cfg.OTLPEndpoint, cfg.AppName, cfg.AppEnv)
	ctx, span := otel.StartTransaction(ctx, "app started")
	defer otel.EndTransaction(span)

	undoLogger, err := logger.NewLogger(cfg)
	if err != nil {
		logger.Error(ctx, "error initializing logger", map[string]interface{}{logger.ErrorKey: err})
		return
	}
	defer undoLogger()

	jwtGen := jwt.NewJWTGenerator(24*time.Hour, cfg.JWTSecretKey)

	pgDb, err := db.NewPostgresDB(db.PostgresDsn{
		Host:     cfg.DbHost,
		User:     cfg.DbUsername,
		Password: cfg.DbPassword,
		Port:     cfg.DbPort,
		Db:       cfg.DbName,
		Env:      cfg.AppEnv,
	})
	if err != nil {
		logger.Error(ctx, "error initializing logger", map[string]interface{}{logger.ErrorKey: err})
		return
	}

	natsMsg := nats.NewNATSConnection(ctx, cfg.NatsURL)
	defer natsMsg.Close()
	natsJS := nats.NewNATSJSConnection(ctx, natsMsg)

	httpClient := http_client.HttpClient()

	srv := pkgHttp.NewHTTPServer(cfg, otel)

	reg := ridesHandler.RegisterHandlerParam{
		Router:     srv.Router.Group(cfg.ApiPrefix + "/v1"),
		DB:         pgDb,
		NatsJS:     natsJS,
		JWTGen:     jwtGen,
		Cfg:        cfg,
		HttpClient: httpClient,
	}
	ridesHandler.RegisterHandler(ctx, reg)

	httpServer := srv.Start(ctx)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	<-quit

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	natsMsg.Close()

	if err := pgDb.Close(); err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	if err := otel.EndAPM(ctx); err != nil {
		logger.Error(ctx, err.Error(), nil)
	}

	logger.Info(ctx, "server is exiting gracefully.", nil)
	_ = logger.Sync()

	if err := otel.EndAPM(ctx); err != nil {
		logger.Error(ctx, err.Error(), nil)
	}
}

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
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/messaging/nats"
	pkgOtel "nebeng-jek/pkg/otel"
	"nebeng-jek/pkg/redis"
	"os"
)

func main() {
	projectEnv := os.Getenv("PROJECT_ENV")
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	cfg := configs.NewConfig(configs.ConfigLoader{
		Env:           projectEnv,
		ConsulAddress: consulAddress,
	}, "./configs/rides")

	err := logger.NewLogger(cfg.AppName, cfg.AppEnv)
	if err != nil {
		logger.Fatal(context.Background(), "error initializing logger", map[string]interface{}{logger.ErrorKey: err})
	}

	otel := pkgOtel.NewOpenTelemetry(cfg.OTLPEndpoint, cfg.AppName, cfg.AppEnv)
	jwtGen := jwt.NewJWTGenerator(24*time.Hour, cfg.JWTSecretKey)

	pgDb, err := db.NewPostgresDB(db.PostgresDsn{
		Host:     cfg.DbHost,
		User:     cfg.DbUsername,
		Password: cfg.DbPassword,
		Port:     cfg.DbPort,
		Db:       cfg.DbName,
	})
	if err != nil {
		logger.Fatal(context.Background(), "error initializing postgres db", map[string]interface{}{logger.ErrorKey: err})
	}

	redisClient := redis.InitConnection(cfg.RedisDB, cfg.RedisHost, cfg.RedisPort,
		cfg.RedisPassword, cfg.RedisAppConfig)

	natsMsg := nats.NewNATSConnection("nats://localhost:4222")
	defer natsMsg.Close()
	natsJS := nats.NewNATSJSConnection(natsMsg)

	srv := pkgHttp.NewHTTPServer(cfg.AppName, cfg.AppEnv, cfg.AppPort, otel)

	reg := ridesHandler.RegisterHandlerParam{
		Router: srv.Router.Group("/v1"),
		Redis:  redisClient,
		DB:     pgDb,
		NatsJS: natsJS,
		JWTGen: jwtGen,
	}
	ridesHandler.RegisterHandler(reg)

	httpServer := srv.Start()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-quit

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Info(ctx, err.Error(), nil)
	}

	if err := pgDb.Close(); err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	if err := otel.EndAPM(); err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	logger.Info(ctx, "server is exiting gracefully.", nil)

	_ = logger.Sync()
}

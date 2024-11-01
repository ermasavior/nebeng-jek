package main

import (
	"context"
	driversHandler "nebeng-jek/internal/drivers/handler"
	"nebeng-jek/pkg/configs"
	pkgHttp "nebeng-jek/pkg/http"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/messaging/nats"
	pkgOtel "nebeng-jek/pkg/otel"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	projectEnv := os.Getenv("PROJECT_ENV")
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	cfg := configs.NewConfig(configs.ConfigLoader{
		Env:           projectEnv,
		ConsulAddress: consulAddress,
	}, "./configs/drivers")

	otel := pkgOtel.NewOpenTelemetry(cfg.OTLPEndpoint, cfg.AppName, cfg.AppEnv)

	undoLogger, err := logger.NewLogger(cfg)
	if err != nil {
		logger.Fatal(ctx, "error initializing logger", map[string]interface{}{logger.ErrorKey: err})
	}
	defer undoLogger()

	natsMsg := nats.NewNATSConnection(ctx, cfg.NatsURL)
	defer natsMsg.Close()
	natsJS := nats.NewNATSJSConnection(ctx, natsMsg)

	jwtGen := jwt.NewJWTGenerator(24*time.Hour, cfg.JWTSecretKey)

	srv := pkgHttp.NewHTTPServer(cfg.AppName, cfg.AppEnv, cfg.AppPort, otel)

	reg := driversHandler.RegisterHandlerParam{
		Router: srv.Router.Group("/v1"),
		NatsJS: natsJS,
		JWTGen: jwtGen,
	}
	driversHandler.RegisterHandler(ctx, reg)

	httpServer := srv.Start(ctx)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	<-quit

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Info(ctx, err.Error(), nil)
	}

	if err := otel.EndAPM(ctx); err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	logger.Info(ctx, "server is exiting gracefully.", nil)

	_ = logger.Sync()
}

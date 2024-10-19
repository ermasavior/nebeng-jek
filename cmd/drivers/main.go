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
	projectEnv := os.Getenv("PROJECT_ENV")
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	cfg := configs.NewConfig(configs.ConfigLoader{
		Env:           projectEnv,
		ConsulAddress: consulAddress,
	}, "./configs/drivers")

	err := logger.NewLogger(cfg.AppName, cfg.AppEnv)
	if err != nil {
		logger.Fatal(context.Background(), "error initializing logger", map[string]interface{}{logger.ErrorKey: err})
	}

	otel := pkgOtel.NewOpenTelemetry(cfg.OTLPEndpoint, cfg.AppName, cfg.AppEnv)

	natsMsg := nats.NewNATSConnection("nats://localhost:4222")
	defer natsMsg.Close()
	natsJS := nats.NewNATSJSConnection(natsMsg)

	jwtGen := jwt.NewJWTGenerator(24*time.Hour, cfg.JWTSecretKey)

	srv := pkgHttp.NewHTTPServer(cfg.AppName, cfg.AppEnv, cfg.AppPort, otel)

	reg := driversHandler.RegisterHandlerParam{
		Router: srv.Router.Group("/v1"),
		NatsJS: natsJS,
		JWTGen: jwtGen,
	}
	driversHandler.RegisterHandler(reg)

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

	if err := otel.EndAPM(); err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	logger.Info(ctx, "server is exiting gracefully.", nil)

	_ = logger.Sync()
}

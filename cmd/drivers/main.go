package main

import (
	"context"
	driversHandler "nebeng-jek/internal/drivers/handler"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/configs"
	pkgHttp "nebeng-jek/pkg/http"
	"nebeng-jek/pkg/logger"
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

	amqpConn, err := amqp.InitAMQPConnection(cfg.AMQPURL)
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp", map[string]interface{}{logger.ErrorKey: err})
	}

	ridesChannel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer ridesChannel.Close()

	srv := pkgHttp.NewHTTPServer(cfg.AppName, cfg.AppEnv, cfg.AppPort, otel)

	driversHandler.RegisterHandler(srv.Router.Group("/"), ridesChannel)

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
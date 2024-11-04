package logger

import (
	"log"
	"nebeng-jek/pkg/configs"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *configs.Config) (func(), error) {
	config := zap.NewDevelopmentConfig()
	if cfg.AppEnv == "production" {
		config = zap.NewProductionConfig()
	}

	// config.OutputPaths = []string{cfg.LogFilePath, "stderr"}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.DisableStacktrace = true

	logZap, err := config.Build()
	if err != nil {
		log.Fatal("error building logger", err)
		return func() {}, err
	}

	logZap = logZap.With(zap.String("app_name", cfg.AppName)).
		WithOptions(zap.WithCaller(true), zap.AddCallerSkip(1))

	logOtel := otelzap.New(logZap,
		otelzap.WithCallerDepth(1), otelzap.WithMinLevel(zapcore.InfoLevel),
		otelzap.WithErrorStatusLevel(zapcore.ErrorLevel),
	)
	undo := otelzap.ReplaceGlobals(logOtel)

	return undo, err
}

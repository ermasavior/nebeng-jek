package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(appName string, env string) error {
	config := zap.NewDevelopmentConfig()
	if env == "production" {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.DisableStacktrace = true

	log, err := config.Build()
	if err != nil {
		return err
	}
	log = log.With(zap.String("app_name", appName))

	logger.zapLogger = log
	return err
}

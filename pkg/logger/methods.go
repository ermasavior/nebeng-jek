package logger

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func Sync() error {
	return otelzap.L().Sync()
}

func Info(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := initBaseLoggerFields(fields)
	otelzap.L().Ctx(ctx).Info(msg, zapFields...)
}

func Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := initBaseLoggerFields(fields)
	otelzap.L().Ctx(ctx).Debug(msg, zapFields...)
}

func Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := initBaseLoggerFields(fields)
	otelzap.L().Ctx(ctx).Warn(msg, zapFields...)
}

func Error(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := initBaseLoggerFields(fields)
	otelzap.L().Ctx(ctx).Error(msg, zapFields...)
}

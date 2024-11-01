package logger

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func Sync() error {
	return otelzap.L().Sync()
}

func Info(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	otelzap.L().Info(msg, zapFields...)
}

func Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	otelzap.L().Debug(msg, zapFields...)
}

func Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	otelzap.L().Warn(msg, zapFields...)
}

func Error(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	otelzap.L().Error(msg, zapFields...)
}

func Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	otelzap.L().Fatal(msg, zapFields...)
}

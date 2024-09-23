package logger

import (
	"context"
)

func Sync() error {
	return logger.zapLogger.Sync()
}

func Info(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	logger.zapLogger.Info(msg, zapFields...)
}

func Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	logger.zapLogger.Debug(msg, zapFields...)
}

func Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	logger.zapLogger.Warn(msg, zapFields...)
}

func Error(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	logger.zapLogger.Error(msg, zapFields...)
}

func Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	zapFields := mapToZapFields(fields)
	zapFields = append(zapFields, initBaseLoggerFields(ctx)...)
	logger.zapLogger.Fatal(msg, zapFields...)
}

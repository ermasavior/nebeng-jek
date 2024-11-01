package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ErrorKey = "error"

func initBaseLoggerFields(fields map[string]interface{}) []zapcore.Field {
	zapFields := []zapcore.Field{}

	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}

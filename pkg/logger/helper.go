package logger

import (
	"nebeng-jek/pkg/utils"
	"context"
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initBaseLoggerFields(ctx context.Context) []zapcore.Field {
	fileName, methodName := traceFileNameAndMethodName()
	fields := []zapcore.Field{
		zap.String("method_name", methodName),
		zap.String("file_name", fileName),
	}

	traceID, ok := ctx.Value(utils.TraceID).(string)
	if ok {
		fields = append(fields, zap.String("trace_id", traceID))
	}

	return fields
}

func traceFileNameAndMethodName() (fileName, methodName string) {
	pc, file, line, _ := runtime.Caller(3)
	details := runtime.FuncForPC(pc)

	return fmt.Sprintf("%s:%d", file, line), details.Name()
}

func mapToZapFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}

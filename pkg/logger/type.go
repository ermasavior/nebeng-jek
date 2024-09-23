package logger

import (
	"go.uber.org/zap"
)

var logger = Logger{
	zapLogger: zap.NewExample(),
}

const ErrorKey = "error"

type Logger struct {
	zapLogger *zap.Logger
}

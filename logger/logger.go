package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the default logger of the application
var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Info is an abstraction to provide an info logger
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug is an abstraction to provide an Debug logger
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error is an abstraction to provide an Error logger
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

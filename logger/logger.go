package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.StacktraceKey = ""

	config.EncoderConfig = encodeConfig

	log, err = config.Build()

	if err != nil {
		panic(err)
	}
}

func LogInfo(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func LogDebug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func LogError(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

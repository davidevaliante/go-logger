package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DefaultEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return encoderConfig
}

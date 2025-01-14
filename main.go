package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options = lumberjack.Logger

var DefaultOptions = &Options{
	Filename:   "app.log",
	MaxSize:    100,
	MaxBackups: 7,
	MaxAge:     7,
	Compress:   true,
}

func New(env string, opts *Options) *zap.Logger {
	if env == "production" {
		return createProductionLogger(opts)
	} else {
		return createDevelopmentLogger(opts)
	}
}

func createProductionLogger(opts *Options) *zap.Logger {
	fileWriteSyncer := zapcore.AddSync(opts)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(jsonEncoder, fileWriteSyncer, zap.InfoLevel)

	logger := zap.New(core)

	return logger
}

func createDevelopmentLogger(opts *Options) *zap.Logger {
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	fileCore := createProductionLogger(opts)
	consoleCore := zapcore.NewCore(jsonEncoder, consoleWriteSyncer, zap.InfoLevel)

	core := zapcore.NewTee(fileCore.Core(), consoleCore)

	logger := zap.New(core)

	return logger
}

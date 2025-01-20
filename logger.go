package logger

import (
	"os"

	"github.com/davidevaliante/constants"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options = lumberjack.Logger
type Logger = zap.Logger

var DefaultOptions = &Options{
	Filename:   "process.log",
	MaxSize:    100,
	MaxBackups: 7,
	MaxAge:     7,
	Compress:   true,
}

func New(env constants.Env, opts *Options) *Logger {
	if env == constants.Development || env == constants.Local {
		return createProductionLogger(opts)
	} else {
		return createDevelopmentLogger(opts)
	}
}

func createProductionLogger(opts *Options) *Logger {
	var fileWriteSyncer zapcore.WriteSyncer
	if opts == nil {
		fileWriteSyncer = zapcore.AddSync(DefaultOptions)
	} else {
		fileWriteSyncer = zapcore.AddSync(opts)
	}

	jsonEncoder := zapcore.NewJSONEncoder(DefaultEncoderConfig())

	core := zapcore.NewCore(jsonEncoder, fileWriteSyncer, zap.InfoLevel)
	logger := zap.New(core)

	return logger
}

func createDevelopmentLogger(opts *Options) *Logger {
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)
	jsonEncoder := zapcore.NewJSONEncoder(DefaultEncoderConfig())

	fileCore := createProductionLogger(opts)
	consoleCore := zapcore.NewCore(jsonEncoder, consoleWriteSyncer, zap.InfoLevel)

	core := zapcore.NewTee(fileCore.Core(), consoleCore)
	logger := zap.New(core)

	return logger
}

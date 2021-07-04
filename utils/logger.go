package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger(logFile string) error {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    500, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})
	consoleWriter := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter),
		zap.InfoLevel,
	)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	return nil
}

func NewLogger(name string) *zap.SugaredLogger {
	return zap.S().Named(name)
}

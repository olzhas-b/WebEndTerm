package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var WorkLogger *zap.Logger

type Config struct {
	EnableConsole bool

	ConsoleJSONFormat bool
	ConsoleLevel      string
}

func Init(isDevelopment bool, cfg *Config) {
	if cfg == nil {
		return
	}
	WorkLogger = initLogger(isDevelopment, cfg)
}

func initLogger(isDevelopment bool, cfg *Config) *zap.Logger {
	var cores []zapcore.Core

	if cfg.EnableConsole {
		level := getZapLevel(cfg.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(isDevelopment, cfg.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	return logger
}

func getZapLevel(level string) zapcore.Level {
	var l zapcore.Level
	err := l.Set(level)
	if err != nil {
		l = zapcore.InfoLevel
	}
	return l
}

func getEncoder(isDevelopment bool, isJSON bool) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if isDevelopment {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

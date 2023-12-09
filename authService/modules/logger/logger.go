package logger

import (
	"context"
	"github.com/mattn/go-colorable"
	"github.com/olzhas-b/PetService/authService/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const InstanceZapLogger int = iota

func InitLogger(cfg *config.Config) {
	var level zapcore.Level
	// if cfg.Debug {
	// 	level = zap.InfoLevel
	// } else {
	// 	level = zap.WarnLevel
	// }
	writer := zapcore.Lock(zapcore.AddSync(colorable.NewColorableStdout()))
	core := zapcore.NewCore(getEncoder(cfg), writer, level)

	newLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(8))
	zap.ReplaceGlobals(newLogger)
}

func getEncoder(cfg *config.Config) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if cfg.Debug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // for color
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	//if isJSON {
	//	return zapcore.NewJSONEncoder(encoderConfig)
	//}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, InstanceZapLogger, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return zap.L()
	}
	if ctxLogger, ok := ctx.Value(InstanceZapLogger).(*zap.Logger); ok {
		return ctxLogger
	} else {
		return zap.L()
	}
}

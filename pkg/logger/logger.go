package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level, environment, fileName string) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	encoder := zap.NewDevelopmentEncoderConfig()
	opts := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeDuration = zapcore.MillisDurationEncoder

	config.OutputPaths = []string{"stderr", fileName}
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig = encoder

	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		config.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		config.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		config.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	log, err := config.Build(opts...)
	if err != nil {
		return nil, err
	}

	return log, nil
}

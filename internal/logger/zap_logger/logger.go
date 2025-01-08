package zaplogger

import (
	"github.com/besean163/gophkeeper/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger() (*Logger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"log"},
		ErrorOutputPaths: []string{"log"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	l, err := config.Build()
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		log: l,
	}

	return logger, nil
}

func convertFields(fields []logger.Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

func (z *Logger) Info(msg string, fields ...logger.Field) {
	z.log.Info(msg, convertFields(fields)...)
}

func (z *Logger) Error(msg string, fields ...logger.Field) {
	z.log.Error(msg, convertFields(fields)...)
}

func (z *Logger) Debug(msg string, fields ...logger.Field) {
	z.log.Debug(msg, convertFields(fields)...)
}

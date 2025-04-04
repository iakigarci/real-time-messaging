package logger

import (
	"sync"

	"real-time-messaging/consumer/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
	once     sync.Once
)

type Logger struct {
	*zap.Logger
}

func New(cfg *config.Config) *Logger {
	once.Do(func() {
		instance = setLoggerFormat(cfg.Logging)
	})
	return &Logger{instance}
}

func setLoggerFormat(cfg config.LogConfig) *zap.Logger {
	config := zap.Config{
		Level:             getLogLevel(cfg.Level),
		Development:       false,
		Encoding:          cfg.Format,
		EncoderConfig:     getEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
	}

	logger, err := config.Build(
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		logger, _ = zap.NewProduction()
	}

	return logger
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getLogLevel(level config.LogLevel) zap.AtomicLevel {
	switch level {
	case config.Debug:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case config.Info:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case config.Trace:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

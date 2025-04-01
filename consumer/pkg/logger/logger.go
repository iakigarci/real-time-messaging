package logger

import (
	"sync"

	"github.com/iakigarci/go-ddd-microservice-template/config"
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

func (l *Logger) ErrorAttrs(msg string, err error, attrs ...map[string]string) {
	fields := []zap.Field{zap.Error(err)}

	if len(attrs) > 0 {
		for key, value := range attrs[0] {
			fields = append(fields, zap.String(key, value))
		}
	}

	l.Logger.Error(msg, fields...)
}

func (l *Logger) InfoAttrs(msg string, attrs ...map[string]string) {
	var fields []zap.Field

	if len(attrs) > 0 {
		fields = make([]zap.Field, 0, len(attrs[0]))
		for key, value := range attrs[0] {
			fields = append(fields, zap.String(key, value))
		}
	}

	l.Logger.Info(msg, fields...)
}

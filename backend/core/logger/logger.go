package logger

import (
	"context"
	"sync"

	"comb.com/banking/internal/config"
)

var (
	instance Logger
	once     sync.Once
)

type Logger interface {
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Warn(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, err error, fields map[string]interface{})
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	Fatal(ctx context.Context, msg string, fields map[string]interface{})
}

type loggerImpl struct{}

func (l *loggerImpl) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	enqueue(ctx, Info, msg, fields, nil)
}

func (l *loggerImpl) Error(ctx context.Context, err error, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["error"] = err.Error()
	enqueue(ctx, Error, err.Error(), fields, err)
}

func (l *loggerImpl) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	enqueue(ctx, Debug, msg, fields, nil)
}

func (l *loggerImpl) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	enqueue(ctx, Warn, msg, fields, nil)
}

func (l *loggerImpl) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	enqueue(ctx, Fatal, msg, fields, nil)
}

// NOTE: Singleton, changed logger setting in config/logger.yaml will take effect after restart
func NewLogger() Logger {
	once.Do(func() {
		cfg := config.LoadConfig()

		var sink Sink
		switch cfg.LogSink {
		case "file":
			fs, err := NewFileSink(cfg.FilePath)
			if err != nil {
				panic(err)
			}
			sink = fs
		case "stdout":
			sink = &StdoutSink{}
		default:
			panic("unknown sink")
		}
		StartDispatcher(sink)
		instance = &loggerImpl{}
	})
	return instance
}

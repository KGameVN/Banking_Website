package logger

import "context"

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

func NewLogger() Logger {
    return &loggerImpl{}
}
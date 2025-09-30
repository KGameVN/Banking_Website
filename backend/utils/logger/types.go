package logger

import "time"

type LogLevel string

const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
	Fatal LogLevel = "FATAL"
)

// LogEntry is canonical structured log message
type LogEntry struct {
	Timestamp time.Time              `json:"ts"`
    Level     LogLevel               `json:"level"`
    Service   string                 `json:"service"`
    Host      string                 `json:"host"`
    Message   string                 `json:"msg"`
    Fields    map[string]interface{} `json:"fields,omitempty"`
    TraceID   string                 `json:"trace_id,omitempty"`
    SpanID    string                 `json:"span_id,omitempty"`
}

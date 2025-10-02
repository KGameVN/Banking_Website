package logger

import (
    "context"
    "os"
    "time"
)

func ingest(ctx context.Context, level LogLevel, msg string, fields map[string]interface{}) LogEntry {
    if fields == nil {
        fields = make(map[string]interface{})
    }

    // Mask PII
    for k, _ := range fields {
        if k == "password" || k == "ssn" {
            fields[k] = "***MASKED***"
        }
    }

    // Enrichment
    host, _ := os.Hostname()

    entry := LogEntry{
        Timestamp: time.Now().UTC(),
        Level:     level,
        Service:   "user-service",
        Host:      host,
        Message:   msg,
        Fields:    fields,
    }

    // Tracing
    attachTrace(ctx, &entry)

    return entry
}

// 4: observability layer: tracing integration
package logger

import (
    "context"

    "go.opentelemetry.io/otel/trace"
)

func attachTrace(ctx context.Context, entry *LogEntry) {
    span := trace.SpanFromContext(ctx)
    sc := span.SpanContext()
    if sc.HasTraceID() {
        entry.TraceID = sc.TraceID().String()
        entry.SpanID = sc.SpanID().String()
    }
}

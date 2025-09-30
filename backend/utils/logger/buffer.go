// 2.processing layer: buffer that holds log entries before they are dispatched

package logger

import "context"

var logQueue = make(chan LogEntry, 10000) // bounded queue

func enqueue(ctx context.Context, level LogLevel, msg string, fields map[string]interface{}, err error) {
    entry := ingest(ctx, level, msg, fields)

    select {
    case logQueue <- entry:
        // enqueued successfully
    default:
        // buffer full → drop + metric
        droppedLogs.Inc()
    }
}

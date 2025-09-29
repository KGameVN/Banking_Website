// 2. processing layer: dispatches log entries from buffer to sinks
package logger

import (
    "time"
)

func StartDispatcher(sink Sink) {
    go func() {
        batch := []LogEntry{}
        ticker := time.NewTicker(1 * time.Second)

        for {
            select {
            case log := <-logQueue:
                batch = append(batch, log)
                if len(batch) >= 100 {
                    sink.Write(batch)
                    batch = nil
                }
            case <-ticker.C:
                if len(batch) > 0 {
                    sink.Write(batch)
                    batch = nil
                }
            }
        }
    }()
}

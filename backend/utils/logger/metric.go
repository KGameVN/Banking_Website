// 4. observability layer: metrics integration
package logger

import "github.com/prometheus/client_golang/prometheus"

var droppedLogs = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "logger_dropped_total",
    Help: "Number of dropped log entries due to full buffer",
})

var bufferSize = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "logger_buffer_size",
    Help: "Current number of log entries in buffer",
})

func init() {
    prometheus.MustRegister(droppedLogs, bufferSize)
}

// 3. output layers: sinks (stdout, file, external systems)

package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"comb.com/banking/config"
	"github.com/segmentio/kafka-go"
)

type Sink interface {
	Write([]LogEntry) error
}

type StdoutSink struct{}

func (s *StdoutSink) Write(entries []LogEntry) error {
	for _, e := range entries {
		b, _ := json.Marshal(e)
		fmt.Println(string(b))
	}
	return nil
}

type FileSink struct {
	file *os.File
}

func NewFileSink(path string) (*FileSink, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileSink{file: f}, nil
}

func (f *FileSink) Write(entries []LogEntry) error {
	for _, e := range entries {
		b, _ := json.Marshal(e)
		f.file.Write(append(b, '\n'))
	}
	return nil
}

type KafkaSink struct {
	writer *kafka.Writer
}

func NewKafkaSink(cfg config.KafkaConfig) *KafkaSink {
	w := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Topic:        cfg.Topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    cfg.BatchSize,
		BatchTimeout: cfg.FlushInterval,
		Async:        true, // non-blocking
	}
	return &KafkaSink{writer: w}
}

func (s *KafkaSink) Write(msg string) error {
	return s.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(msg),
		},
	)
}

func (s *KafkaSink) Close() error {
	return s.writer.Close()
}

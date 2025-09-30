package config

import "time"

type KafkaConfig struct {
	Brokers       []string
	Topic         string
	BatchSize     int
	FlushInterval time.Duration
}

func GetKafka() KafkaConfig {
	return KafkaConfig{
		Brokers:       []string{"localhost:9092"}, // thay bằng danh sách broker Kafka của bạn
		Topic:         "app-logs",             // tên topic muốn đẩy log
		BatchSize:     200,                    // số message gom trong 1 batch
		FlushInterval: 250 * time.Millisecond, // thời gian flush batch
	}
}

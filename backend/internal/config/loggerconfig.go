package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Service         string
	LogSink         string
	FilePath        string
	WALPath         string
	BatchSize       int
	WorkerCount     int
	KafkaBrokers    []string
	KafkaTopic      string
	TransactionalID string
}

func LoadConfig() Config {
	viper.SetConfigName("logger") // tên file cấu hình logger.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./yaml")
	viper.AutomaticEnv() // allow ENV overrides

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return Config{
		Service:         viper.GetString("app.service"),
		LogSink:         viper.GetString("log.sink"),
		FilePath:        viper.GetString("log.file_path"),
		WALPath:         viper.GetString("log.wal_path"),
		BatchSize:       viper.GetInt("log.batch_size"),
		WorkerCount:     viper.GetInt("log.worker_count"),
		KafkaBrokers:    viper.GetStringSlice("kafka.brokers"),
		KafkaTopic:      viper.GetString("kafka.topic"),
		TransactionalID: viper.GetString("kafka.transactional_id"),
	}
}

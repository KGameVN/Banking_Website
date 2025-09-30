package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Level     string
    Sink      string
    FilePath  string
    BatchSize int
}

func LoadConfig() Config {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    return Config{
        Level:     viper.GetString("log.level"),
        Sink:      viper.GetString("log.sink"),
        FilePath:  viper.GetString("log.file_path"),
        BatchSize: viper.GetInt("log.batch_size"),
    }
}

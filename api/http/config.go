package server

import "github.com/spf13/viper"

type Config struct {
	// API server port
	Port int
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Port: viper.GetInt("port"),
	}

	if config.Port == 0 {
		config.Port = 8080
	}

	return config, nil
}

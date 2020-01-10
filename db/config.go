package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	dbName string
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbType string
	debug  bool
}

func LoadConfig() *Config {
	config := &Config{
		dbName: viper.GetString("dbName"),
		dbHost: viper.GetString("dbHost"),
		dbPort: viper.GetString("dbPort"),
		dbUser: viper.GetString("dbUser"),
		dbPass: viper.GetString("dbPass"),
		dbType: viper.GetString("dbType"),
		debug:  viper.GetBool("debug"),
	}
	return config
}

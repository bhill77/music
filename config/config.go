package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	StoragePath string
	DB          DB
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		Port:        viper.GetString("port"),
		StoragePath: viper.GetString("storage_path"),
		DB: DB{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Database: viper.GetString("db.database"),
		},
	}
}

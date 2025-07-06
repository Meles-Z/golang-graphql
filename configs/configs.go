package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func LoadConfig() (*Config, error) {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.BindEnv()

	cfg := Config{
		DBConfig{
			Host:     viper.GetString(""),
			Port:     viper.GetInt(""),
			DBName:   viper.GetString(""),
			User:     viper.GetString(""),
			Password: viper.GetString(""),
		},
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore or log as needed
			log.Fatalf("Error load config file: %v", err)
		} else {
			// Config file was found but another error occurred
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	return &cfg, nil
}

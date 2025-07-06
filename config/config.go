package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Care all our configuration
type Config struct {
	DB   DatabaseConfig
	Auth AuthConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error to reading env file:", err)
			return nil, err
		}
		fmt.Println("Error to read config file:", err)
	}

	cfg := Config{
		DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		AuthConfig{
			Secret: viper.GetString("WEB_SECRET"),
		},
	}
	return &cfg, nil

}

// config/config.go
package config

import (
	"github.com/spf13/viper"
)

// Config represents the server configuration structure
type Config struct {
	Server struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"server"`
}

// LoadConfig loads configuration from config.yaml file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

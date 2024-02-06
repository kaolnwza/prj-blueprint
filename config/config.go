package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConf DatabaseConfig `mapstructure:"db"`
}

type (
	DatabaseConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	}
)

func NewConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Config file not found: %w", err))
		}

		panic(fmt.Errorf("Config read failed: %w", err))
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("Config unmarshal failed: %w", err))
	}

	return conf
}

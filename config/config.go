package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	VERSION = "1.0.0" // to be in sync with git tag
)

type Config struct {
	Env              string             `mapstructure:"env"`
	DatabaseConf     DatabaseConfig     `mapstructure:"db"`
	RedisConf        RedisConfig        `mapstructure:"redis"`
	LogConf          LogConfig          `mapstructure:"log"`
	HttpConf         HttpConfig         `mapstructure:"http"`
	MicroserviceConf MicroserviceConfig `mapstructure:"microservice"`
	ExternalApiConf  ExternalApiConfig  `mapstructure:"external_api"`
}

type HttpConfig struct {
	MaxIdleConns        int `mapstructure:"max_idle_conns"`
	MaxConnsPerHost     int `mapstructure:"max_conns_per_host"`
	MaxIdleConnsPerHost int `mapstructure:"max_idle_conns_per_host"`
}

type LogConfig struct {
	Level           string `mapstructure:"level"`
	JsonFormat      bool   `mapstructure:"json_format"`
	ShowBody        bool   `mapstructure:"show_body"`
	ShowBodyOnError bool   `mapstructure:"show_body_on_error"`
	Pretty          bool   `mapstructure:"pretty"`
	IncludeHealth   bool   `mapstructure:"include_health"`
}

func New() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("[Config] file not found: %w", err))
		}

		panic(fmt.Errorf("[Config] read failed: %w", err))
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("[Config] unmarshal failed: %w", err))
	}

	conf.parseSecret()
	if err := conf.setupHttpPoolClient(); err != nil {
		panic(fmt.Errorf("[Config]failed to setHttpPoolClients, err = %w", err))
	}

	return conf
}

func (c *Config) setupHttpPoolClient() error {
	c.MicroserviceConf.setupHttpPoolClient(c.HttpConf)
	return c.ExternalApiConf.setupHttpPoolClients(c.HttpConf)
}

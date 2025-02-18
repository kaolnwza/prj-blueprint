package config

import (
	"log"
	"net/http"
	"time"
)

type EndpointConf struct {
	Method string `mapstructure:"http_method"`
	Url    string `mapstructure:"endpoint"`
}

type BaseMicrosvcConfig[T any] struct {
	BaseUrl   string `mapstructure:"base_url"`
	Endpoints T      `mapstructure:"endpoints"`
}

type MicroserviceConfig struct {
	Timeout        string                                `mapstructure:"timeout"`
	CustomerConf   BaseMicrosvcConfig[CustomerEndpoints] `mapstructure:"customer_service"`
	httpPoolClient http.Client
}

type (
	CustomerEndpoints struct {
		InqByCif EndpointConf `mapstructure:"inq_by_cif"`
	}
)

func (c MicroserviceConfig) GetHttpClient() http.Client {
	return c.httpPoolClient
}

func (c *MicroserviceConfig) setupHttpPoolClient(httpConf HttpConfig) {
	defaultTimeout := time.Second * 30
	exp, err := time.ParseDuration(c.Timeout)
	if err != nil {
		log.Println("[Config] failed to parse timeout on microservice http client -> parsed default timeout, err = %w", err)
		exp = defaultTimeout
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = httpConf.MaxIdleConns
	t.MaxConnsPerHost = httpConf.MaxConnsPerHost
	t.MaxIdleConnsPerHost = httpConf.MaxIdleConnsPerHost

	c.httpPoolClient = http.Client{Transport: t, Timeout: exp}
}

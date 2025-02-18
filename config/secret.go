package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	secretKub = "SECRET_KUB"
)

func (c *Config) parseSecret() {
	vi := viper.New()
	vi.SetEnvPrefix("secret")
	vi.AutomaticEnv()
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if v := vi.GetString(secretKub); v != "" {
		c.ExternalApiConf.UserCenterConf.CertsConf.Certs = v
	}
}

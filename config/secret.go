package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	secretUserCenterCerts = "GU_SECRET"
)

func parseSecret(conf *Config) {
	vi := viper.New()
	vi.SetEnvPrefix("secret")
	vi.AutomaticEnv()
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if v := vi.GetString(secretUserCenterCerts); v != "" {
		conf.ExternalApiConf.UserCenterConf.CertsConf.Certs = v
	}
}

func getSecret(vi *viper.Viper, envKey string) string {
	if v := vi.GetString(envKey); v != "" {
		return v
	}

	return ""
}

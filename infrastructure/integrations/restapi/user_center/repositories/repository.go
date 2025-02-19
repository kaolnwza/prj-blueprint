package repositories

import (
	"github.com/kaolnwza/proj-blueprint/config"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/ports"

	"github.com/kaolnwza/proj-blueprint/libs/api"
)

type repo struct {
	httpConf   config.HttpConfig
	httpCli    api.HttpClient
	extApiConf config.ExternalApiConfig
}

func New(
	httpConf config.HttpConfig,
	httpCli api.HttpClient,
	extApiConf config.ExternalApiConfig,
) ports.Repository {
	return repo{}
}

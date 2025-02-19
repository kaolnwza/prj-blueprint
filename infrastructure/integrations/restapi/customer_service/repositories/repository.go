package repositories

import (
	"github.com/kaolnwza/proj-blueprint/config"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/ports"

	"github.com/kaolnwza/proj-blueprint/libs/api"
)

type repo struct {
	httpConf     config.HttpConfig
	httpCli      api.HttpClient
	microsvcConf config.MicroserviceConfig
}

func New(
	httpConf config.HttpConfig,
	httpCli api.HttpClient,
	microsvcConf config.MicroserviceConfig,
) ports.Repository {
	return repo{}
}

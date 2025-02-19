package service

import (
	custSvcPorts "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/ports"
	userCenterPorts "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/ports"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/ports"
)

type service struct {
	userRepo       ports.Repository
	custSvcRepo    custSvcPorts.Repository
	userCenterRepo userCenterPorts.Repository
}

func New(
	userRepo ports.Repository,
	custSvcRepo custSvcPorts.Repository,
	userCenterRepo userCenterPorts.Repository,
) ports.Service {
	return service{
		userRepo:       userRepo,
		custSvcRepo:    custSvcRepo,
		userCenterRepo: userCenterRepo,
	}
}

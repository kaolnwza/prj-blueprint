package handlers

import "github.com/kaolnwza/proj-blueprint/internal/domains/user/core/ports"

type handler struct {
	userSvc ports.Service
}

func New(userSvc ports.Service) handler {
	return handler{
		userSvc: userSvc,
	}
}

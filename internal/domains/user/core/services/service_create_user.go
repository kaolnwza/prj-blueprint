package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/response"
)

func (s service) CreateUser(ctx context.Context, req models.ReqUser) error {
	user, err := s.userRepo.GetUserById(ctx, uuid.New())
	if err != nil {
		return response.NewInternalError(response.InternalServerError, response.MsgInternalServerError, err)
	}

	if user == nil {
		return response.NewInternalError(response.UserNotFound, response.MsgUserNotFound, fmt.Errorf(response.MsgUserNotFound))
	}

	if err := s.userRepo.CreateUser(ctx, req); err != nil {
		return response.NewInternalError(response.InternalServerError, response.MsgInternalServerError, err)
	}

	return nil
}

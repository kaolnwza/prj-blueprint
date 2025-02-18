package ports

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/internal/domains/exam-user/core/models"
)

type Service interface {
	CreateUser(ctx context.Context, user models.ReqUser) error
	GetUserById(ctx context.Context, userId string) (models.RespUser, error)
}

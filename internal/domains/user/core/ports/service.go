package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
)

type Service interface {
	CreateUser(ctx context.Context, user models.ReqUser) error
	InqUser(ctx context.Context, userId uuid.UUID) (models.RespUser, error)
}

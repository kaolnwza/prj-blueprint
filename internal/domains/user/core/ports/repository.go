package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.ReqUser) error
	GetUserById(ctx context.Context, userId uuid.UUID) (*entities.Users, error)
}

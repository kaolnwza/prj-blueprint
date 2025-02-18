package ports

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/internal/domains/exam-user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.ReqUser) error
	GetUserById(ctx context.Context, userId string) (entities.Users, error)
}

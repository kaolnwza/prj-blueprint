package repositories

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/entities"
)

func (r repo) CreateUser(ctx context.Context, user models.ReqUser) error {
	ent := entities.Users{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	return r.gormDb.New(ctx).Create(&ent).Error
}

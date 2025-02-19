package service

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/entities"
	"go.uber.org/mock/gomock"
)

func (t *TestSuite) TestCreateUser_Some_Behav_ReturnSuccess() {
	ctx := context.Background()
	inq := models.ReqUser{}

	t.mUserRepo.EXPECT().GetUserById(ctx, gomock.Any()).Return(&entities.Users{}, nil)

	t.mUserRepo.EXPECT().CreateUser(ctx, inq).Return(nil)

	actual := t.s.CreateUser(ctx, inq)
	t.NoError(actual)
}

package service

import (
	"context"

	custSvcModels "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/models"
	userCtModels "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/models"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/api"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func (t *TestSuite) TestInqUser_Some_Behav_ReturnSuccessAndUser() {
	inqUserId := uuid.New()
	ctx := context.Background()

	t.mCustSvcRepo.EXPECT().ExamMicrosvcInqUserKub(ctx, gomock.Any()).Return(
		api.BaseResponse[custSvcModels.RespInqUser]{
			Code: "0",
			Data: custSvcModels.RespInqUser{
				Name: "te",
			},
		}, nil,
	)

	t.mUserCtRepo.EXPECT().ExamExternalApiInqUserKub(ctx, gomock.Any()).Return(
		api.BaseResponse[userCtModels.RespInqUser]{
			Code: "0",
			Data: userCtModels.RespInqUser{
				Name: "st",
			},
		}, nil,
	)

	expect := models.RespUser{
		Firstname: "te",
		Lastname:  "st",
	}

	actual, err := t.s.InqUser(ctx, inqUserId)
	t.NoError(err)
	t.Equal(expect, actual)
}

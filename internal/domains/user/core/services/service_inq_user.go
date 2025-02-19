package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	custSvcExct "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/exceptions"
	custSvcModels "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/models"
	userCtModels "github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/models"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/response"
)

func (s service) InqUser(ctx context.Context, userId uuid.UUID) (models.RespUser, error) {
	reqInqUser := custSvcModels.ReqInqUser{
		Id: userId.String(),
	}

	cust, err := s.custSvcRepo.ExamMicrosvcInqUserKub(ctx, reqInqUser)
	if err != nil {
		return models.RespUser{}, response.NewInternalError(response.ExternalServiceUnavailable, response.MsgExternalServiceUnavailable, err)
	}

	if cust.Code == custSvcExct.UserNotFound {
		return models.RespUser{}, response.NewNotFoundError(response.UserNotFound, response.MsgUserNotFound, fmt.Errorf(response.MsgUserNotFound))
	}

	reqUserCt := userCtModels.ReqInqUser{
		Id: userId.String(),
	}

	ucInfo, err := s.userCenterRepo.ExamExternalApiInqUserKub(ctx, reqUserCt)
	if err != nil {
		return models.RespUser{}, response.NewInternalError(response.ExternalServiceUnavailable, response.MsgExternalServiceUnavailable, err)
	}

	resp := models.RespUser{
		Firstname: cust.Data.Name,
		Lastname:  ucInfo.Data.Name,
	}

	return resp, nil
}

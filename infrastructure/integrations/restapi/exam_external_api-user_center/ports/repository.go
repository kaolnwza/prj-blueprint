package ports

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/exam_external_api-user_center/models"
	"github.com/kaolnwza/proj-blueprint/libs/api"
)

type Repository interface {
	InqUserKub(ctx context.Context, req models.ReqInqUser) (api.BaseResponse[models.RespInqUser], error)
}

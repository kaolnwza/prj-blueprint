package ports

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/models"
	"github.com/kaolnwza/proj-blueprint/libs/api"
)

type Repository interface {
	ExamExternalApiInqUserKub(ctx context.Context, req models.ReqInqUser) (api.BaseResponse[models.RespInqUser], error)
}

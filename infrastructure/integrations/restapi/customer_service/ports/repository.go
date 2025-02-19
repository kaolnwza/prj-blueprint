package ports

import (
	"context"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/customer_service/models"
	"github.com/kaolnwza/proj-blueprint/libs/api"
)

type Repository interface {
	ExamMicrosvcInqUserKub(ctx context.Context, req models.ReqInqUser) (api.BaseResponse[models.RespInqUser], error)
}

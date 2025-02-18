package repositories

import (
	"context"
	"net/http"

	"github.com/kaolnwza/proj-blueprint/config"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/exam_external_api-user_center/models"
	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/ports"

	"github.com/kaolnwza/proj-blueprint/libs/api"
	"github.com/kaolnwza/proj-blueprint/libs/constants"
	"github.com/kaolnwza/proj-blueprint/libs/utils"
)

type repo struct {
	httpConf     config.HttpConfig
	httpCli      api.HttpClient
	microsvcConf config.MicroserviceConfig
}

func New(
	httpConf config.HttpConfig,
	httpCli api.HttpClient,
	microsvcConf config.MicroserviceConfig,
) ports.Repository {
	return repo{}
}

func (r repo) ExamMicrosvcInqUserKub(ctx context.Context, req models.ReqInqUser) (api.BaseResponse[models.RespInqUser], error) {
	conf := r.microsvcConf.CustomerConf
	epConf := conf.Endpoints.InqByCif
	cli := r.microsvcConf.GetHttpClient()
	url, err := r.httpCli.BuildEndpoint(api.ApiUrl{
		BaseUrl:  conf.BaseUrl,
		Endpoint: epConf.Url,
	})
	if err != nil {
		return api.BaseResponse[models.RespInqUser]{}, nil
	}

	body, err := api.SerializeObject(req)
	if err != nil {
		return api.BaseResponse[models.RespInqUser]{}, nil
	}

	header := utils.NewHeader(ctx)
	header.Set(constants.ContentType, constants.ApplicationJson)

	reqBody := http.Request{
		Method: epConf.Method,
		URL:    url,
		Body:   body,
		Header: header,
	}

	resp := api.BaseResponse[models.RespInqUser]{}
	return resp, r.httpCli.NewRequest(ctx, cli, &reqBody, &resp)
}

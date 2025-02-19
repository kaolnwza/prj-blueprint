package repositories

import (
	"context"
	"net/http"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/models"
	"github.com/kaolnwza/proj-blueprint/libs/api"
	"github.com/kaolnwza/proj-blueprint/libs/consts"
	"github.com/kaolnwza/proj-blueprint/libs/utils"
)

func (r repo) ExamExternalApiInqUserKub(ctx context.Context, req models.ReqInqUser) (api.BaseResponse[models.RespInqUser], error) {
	conf := r.extApiConf.UserCenterConf
	epConf := conf.Endpoints.Inq
	cli := conf.GetHttpClient()
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
	header.Set(consts.ContentType, consts.ApplicationJson)

	reqBody := http.Request{
		Method: epConf.Method,
		URL:    url,
		Body:   body,
		Header: header,
	}

	resp := api.BaseResponse[models.RespInqUser]{}
	return resp, r.httpCli.NewRequest(ctx, cli, &reqBody, &resp)
}

package repositories

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/ports"
	"github.com/kaolnwza/proj-blueprint/libs/api"
	"github.com/kaolnwza/proj-blueprint/libs/consts"
	"github.com/kaolnwza/proj-blueprint/libs/utils"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	mCtx context.Context
}

/*
	cmd
	go test github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/repositories -v

	=== Run Specific Test
	go test ... -testify.m TestName

*/

func (t *TestSuite) SetupTest() {
	t.mCtx = utils.SetContext(context.Background(), consts.CtxHeaderKey, http.Header{})
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) NewServer(status int, jsonStr string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(jsonStr))
	}))
}

func (t *TestSuite) NewRepo(url string) ports.Repository {
	extConf := config.ExternalApiConfig{
		UserCenterConf: config.BaseExtApiConf[config.UserCenterEndpoints]{
			BaseUrl: url,
			Timeout: "20s",
			Endpoints: config.UserCenterEndpoints{
				Inq: config.EndpointConf{
					Method: http.MethodPost,
					Url:    "/",
				},
			},
		},
	}

	httpCli := api.New(config.LogConfig{}, config.HttpConfig{})
	return New(config.HttpConfig{}, httpCli, extConf)
}

// 	// extConf := infrastructure.ExternalApiConfig{
// 	// 	NextPartnerAuthPin: infrastructure.BaseExternalApiConfig[infrastructure.NextPartnerAuthPinListConf]{
// 	// 		BaseUrl: url,
// 	// 		Endpoints: infrastructure.NextPartnerAuthPinListConf{
// 	// 			PinVerifyStatus: infrastructure.EndpointConf{
// 	// 				// EndpointMethod: http.MethodPost,
// 	// 				// EndpointUrl:    "/",
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }

// 	return New(api.New(api.NewHttp("", "5s", infrastructure.HttpConnectionPool{}), infrastructure.LogConfig{}, extConf))
// }

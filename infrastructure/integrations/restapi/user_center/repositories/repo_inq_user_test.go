package repositories

import (
	"log"
	"net/http"

	"github.com/kaolnwza/proj-blueprint/infrastructure/integrations/restapi/user_center/models"
)

func (t *TestSuite) TestInqUser_Some_Behav_ReturnSuccess() {
	jsonString := `{
		"code": 123,
		"message": "gutest",
		"data": {
			"id": "123",
			"name": "ab"
		}
 }`

	srv := t.NewServer(http.StatusOK, jsonString)
	defer srv.Close()

	// resp, err := t.NewRepo(srv.URL).PinVerifyStatus(t.mCtx, models.ReqPinVerifyStatus{})
	resp, err := t.NewRepo(srv.URL).ExamExternalApiInqUserKub(t.mCtx, models.ReqInqUser{})
	t.NoError(err)
	_ = resp
	log.Println("resp", resp)
}

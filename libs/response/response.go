package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaolnwza/proj-blueprint/pkg/logger"
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResponse(message string, data interface{}) response {
	resp := response{
		Code:    Success,
		Message: message,
		Data:    data,
	}
	return resp
}

func MakeSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BuildResponse(MsgSuccess, data))
}

func MakeError(c *gin.Context, err error) {
	logger := logger.GetByContext(c.Request.Context())
	template := errorResponseBuilder(err)
	logger.Error(template.Error())

	c.AbortWithStatusJSON(template.StatusCode, buildErrorResponse(template.Message, template.Code))
}

func buildErrorResponse(message string, code string) response {
	errorResp := response{
		Code:    code,
		Message: message,
	}

	return errorResp
}

func errorResponseBuilder(err error) errorTemplate {
	switch err.(type) {
	case errorTemplate:
		e := err.(errorTemplate)
		return errorTemplate{
			StatusCode: e.StatusCode,
			Code:       e.Code,
			Message:    e.Message,
			Err:        e.Err,
		}
	}

	return errorTemplate{
		StatusCode: 500,
		Code:       "999",
		Message:    "Unknown",
		Err:        err,
	}
}

func NewInternalError(code string, message string, err error) error {
	return errorTemplate{
		StatusCode: http.StatusInternalServerError,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func NewNotFoundError(code string, message string, err error) error {
	return errorTemplate{
		StatusCode: http.StatusNotFound,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func NewConflictError(code string, message string, err error) error {
	return errorTemplate{
		StatusCode: http.StatusConflict,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func NewBadRequestError(code string, message string, err error) error {
	return errorTemplate{
		StatusCode: http.StatusBadRequest,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

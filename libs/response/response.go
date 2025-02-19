package response

import "net/http"

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponseBuilder(err error) errorTemplate {
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

func NewInternalError(code errCode, message errMsg, err error) error {
	return errorTemplate{
		StatusCode: http.StatusInternalServerError,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

func NewBadRequestError(code errCode, message errMsg, err error) error {
	return errorTemplate{
		StatusCode: http.StatusBadRequest,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

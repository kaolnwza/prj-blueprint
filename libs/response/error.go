package response

type (
	errCode string
	errMsg  string
)

type errorTemplate struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e errorTemplate) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

const (
	Success                    = "0"
	UserNotFound               = "405"
	InternalServerError        = "500"
	ExternalServiceUnavailable = "700"
)

const (
	MsgSuccess                    = "success."
	MsgInternalServerError        = "internal server error."
	MsgUserNotFound               = "user not found."
	MsgExternalServiceUnavailable = "external service is unavailable."
)

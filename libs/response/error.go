package response

type (
	errCode string
	errMsg  string
)

type errorTemplate struct {
	StatusCode int     `json:"-"`
	Code       errCode `json:"code"`
	Message    errMsg  `json:"message"`
	Err        error   `json:"-"`
}

func (e errorTemplate) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

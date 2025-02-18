package models

type (
	ReqUser struct {
		Firstname string `json:"firstname" validate:"required,min=1,max=255"`
		Lastname  string `json:"lastname" validate:"required,min=1,max=255"`
	}

	RespUser struct {
		Id        int    `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
)

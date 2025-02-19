package models

type ReqInqUser struct {
	Id string `json:"id"`
}

type RespInqUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

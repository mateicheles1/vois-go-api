package models

type RequestBodyList struct {
	Owner string   `json:"owner" binding:"required"`
	Todos []string `json:"todos" binding:"required"`
}

type ResponseBodyList struct {
	Id    string `json:"listid,omitempty"`
	Owner string `json:"owner"`
	Todos []ToDo `json:"todos"`
}

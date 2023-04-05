package models

type RequestBodyList struct {
	Owner string   `json:"owner" binding:"required"`
	Todos []string `json:"todos" binding:"required"`
}

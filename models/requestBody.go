package models

type RequestBodyList struct {
	Owner string   `json:"owner"`
	Todos []string `json:"todos" binding:"required"`
}
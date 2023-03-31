package models

type ToDo struct {
	Id      string `json:"todoid,omitempty"`
	Content string `json:"content" binding:"required"`
}

type ToDoList struct {
	Id    string           `json:"listid,omitempty"`
	Owner string           `json:"owner" binding:"required"`
	Todos map[string]*ToDo `json:"todos,omitempty"`
}

type RequestBodyList struct {
	Owner string   `json:"owner" binding:"required"`
	Todos []string `json:"todos" binding:"required"`
}

var Data = make(map[string]*ToDoList)
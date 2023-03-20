package models

type ToDo struct {
	Id      string `json:"todoid,omitempty"`
	Listid  string `json:"todolistid,omitempty"`
	Content string `json:"content" binding:"required"`
}

type ToDoList struct {
	Id    string           `json:"listid,omitempty"`
	Owner string           `json:"owner" binding:"required"`
	Todos map[string]*ToDo `json:"todos,omitempty"`
}

var Data = make(map[string]*ToDoList)
package models

type ToDo struct {
	Id      string `json:"todoid,omitempty"`
	Listid  string `json:"todolistid,omitempty"`
	Content string `json:"content,omitempty"`
}

type ToDoList struct {
	Id    string           `json:"listid,omitempty"`
	Owner string           `json:"owner,omitempty"`
	Todos map[string]*ToDo `json:"todos,omitempty"`
}

var Data = make(map[string]*ToDoList)
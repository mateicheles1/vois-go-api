package models

type ToDo struct {
	Listid  string `json:"todolistid,omitempty"`
	Content string `json:"content,omitempty"`
}

type ToDoList struct {
	Owner string           `json:"owner,omitempty"`
	Todos map[string]*ToDo `json:"todos,omitempty"`
}



var Data = make(map[string]*ToDoList)
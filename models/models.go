package models

type ToDo struct {
	Id        string `json:"todoid,omitempty"`
	ListId    string `json:"todolistid,omitempty"`
	Content   string `json:"content,omitempty"`
	Completed bool   `json:"completed"`
}

type ToDoList struct {
	Id    string           `json:"listid,omitempty"`
	Owner string           `json:"owner" binding:"required"`
	Todos map[string]*ToDo `json:"todos,omitempty"`
}

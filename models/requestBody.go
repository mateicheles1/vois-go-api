package models

type RequestBodyList struct {
	Id    string   `json:"listid,omitempty"`
	Owner string   `json:"owner" binding:"required"`
	Todos []string `json:"todos"`
}

type ToDoList struct {
	Id    string  `gorm:"type:uuid;primary_key" json:"listId,omitempty"`
	Owner string  `gorm:"not null" json:"owner,omitempty"`
	Todos []*ToDo `gorm:"foreignKey:ListId;constraint:OnDelete:CASCADE" json:"todos"`
}

type ToDo struct {
	Id        string `gorm:"type:uuid;primary_key" json:"todoid,omitempty"`
	ListId    string `gorm:"type:uuid;not null" json:"todoListId,omitempty"`
	Content   string `gorm:"not null" json:"content,omitempty"`
	Completed bool   `gorm:"not null" json:"completed"`
}

package models

type User struct {
	Username string      `gorm:"primary_key" json:"username" binding:"required"`
	Password string      `gorm:"not null" json:"password" binding:"required"`
	Role     string      `gorm:"not null" json:"role"`
	Lists    []*ToDoList `gorm:"foreignKey:Owner;constraint:OnDelete:Cascade" json:"lists,omitempty"`
}

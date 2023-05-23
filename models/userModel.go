package models

type User struct {
	Username string `gorm:"not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Role     string `gorm:"not null" json:"role"`
}

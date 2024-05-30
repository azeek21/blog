package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string    `json:"fullname"`
	Username string    `json:"username"`
	Email    string    `gorm:"index;unique" json:"email"`
	Articles []Article `gorm:"foreignKey:AuthorID"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

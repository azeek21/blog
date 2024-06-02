package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `gorm:"index;unique" json:"email"`
	Password string
	Articles []Article `gorm:"foreignKey:AuthorID" json:"articles"`
	Role     Role      `gorm:"foreignKey:RoleCode;references:Code"`
	RoleCode string    `json:"role"`
}

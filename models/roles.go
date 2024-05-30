package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code        string  `gorm:"index;unique" json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Users       []*User `gorm:"many2many:user_roles;"`
}

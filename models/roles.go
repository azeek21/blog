package models

import "gorm.io/gorm"

const ROLE_MODEL_NAME = "role"

type Role struct {
	gorm.Model
	Code        string `gorm:"index;unique" json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	AuthorID uint    `json:"authorId"`
	ImageUrl *string `json:"imageUrl"`
}

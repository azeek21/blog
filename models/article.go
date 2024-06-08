package models

import "gorm.io/gorm"

const ARTICLE_MODEL_NAME = "article"

type Article struct {
	gorm.Model
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	AuthorID uint    `json:"authorId"`
	ImageUrl *string `json:"imageUrl"`
}

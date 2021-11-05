package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content  string `json:"content" binding:"required,min=4,max=1024"`
	AuthorID uint   `json:"user_id"`
}

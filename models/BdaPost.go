package models

import "gorm.io/gorm"

type BdaPost struct {
	gorm.Model
	AuthorID User   `json:"id"`
	Content  string `json:"content" binding:"required,min=4,max=1024"`
}

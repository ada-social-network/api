package models

import "gorm.io/gorm"

type BdaPost struct {
	gorm.Model
	Title   string `json:"title" binding:"required,min=4,max=100"`
	Content string `json:"content" binding:"required,min=4,max=1024"`
	UserID  uint   `json:"user_id"`
}

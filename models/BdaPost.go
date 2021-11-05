package models

import "gorm.io/gorm"

type BdaPost struct {
	gorm.Model
	Content string `json:"content" binding:"required,min=4,max=1024"`
	UserID  uint   `json:"user_id"`
}

package models

import "gorm.io/gorm"

// BdaPost define a bda post resource
type BdaPost struct {
	gorm.Model
	Title   string `json:"title" binding:"required,min=4,max=100"`
	Content string `json:"content" binding:"required,min=4,max=1024"`

	// By default, gorm will try to use UserID as a foreign key to the model User
	UserID uint `json:"user_id" binding:"required"`
}
package models

import "gorm.io/gorm"

// Comment define comment for a post
type Comment struct {
	gorm.Model
	UserID    uint   `json:"user_id" binding:"required"`
	BdapostID uint   `json:"bdapost_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

// we do not have comment for all posts for now only for BDA

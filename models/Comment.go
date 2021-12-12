package models

import uuid "github.com/satori/go.uuid"

// Comment define comment for a post
type Comment struct {
	Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"user_id" `
	BdapostID uuid.UUID `gorm:"type=uuid" json:"bdapost_id" binding:"required"`
	Content   string    `json:"content" binding:"required"`
}

// we do not have comment for all posts for now only for BDA

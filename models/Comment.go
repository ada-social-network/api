package models

import uuid "github.com/satori/go.uuid"

// Comment define comment for a post
type Comment struct {
	Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	BdaPostID uuid.UUID `gorm:"type=uuid" json:"bdapostId" binding:"required"`
	Content   string    `json:"content" binding:"required"`
}

// we do not have comment for all posts for now only for BDA

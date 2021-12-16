package models

import uuid "github.com/satori/go.uuid"

// Post define post resource
type Post struct {
	Base
	Title   string `json:"title" binding:"required,min=4,max=100"`
	Content string `json:"content" binding:"required,min=4,max=21474"`
	// By default, gorm will try to use UserID as a foreign key to the model User
	UserID uuid.UUID `gorm:"type=uuid" json:"user_id"`
}

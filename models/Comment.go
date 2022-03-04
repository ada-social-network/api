package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Comment define comment for a bda post
type Comment struct {
	Base
	UserID    uuid.UUID `gorm:"type=uuid" `
	BdaPostID uuid.UUID `gorm:"type=uuid"`
	Content   string
	Likes     []Like
}

// CommentResponse define the response for a comment on a bda post
type CommentResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userID"`
	BdaPostID uuid.UUID `json:"bdaPostID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CommentRequestURI define the request uri for comment
type CommentRequestURI struct {
	BdaPostID string `uri:"id" binding:"uuid"`
}

// CommentRequest define the request for comment
type CommentRequest struct {
	Content string `json:"content" binding:"required,min=4,max=1024"`
}

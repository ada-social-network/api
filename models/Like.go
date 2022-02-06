package models

import uuid "github.com/satori/go.uuid"

// Like define informations about a like
type Like struct {
	Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	BdaPostID uuid.UUID `gorm:"type=uuid" json:"bdaPostId,omitempty"`
	PostID    uuid.UUID `gorm:"type=uuid" json:"postId,omitempty"`
	CommentID uuid.UUID `gorm:"type=uuid" json:"commentId,omitempty"`
}

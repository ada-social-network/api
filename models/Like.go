package models

import uuid "github.com/satori/go.uuid"

// Like define informations about a like
type Like struct {
	Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	BdaPostID uuid.UUID `gorm:"type=uuid" json:"bdapostId"`
	PostID    uuid.UUID `gorm:"type=uuid" json:"postId"`
	CommentID uuid.UUID `gorm:"type=uuid" json:"commentId"`
}

// CountAndLikes defines the count of likes and the likes associated to a resource
type CountAndLikes struct {
	Likes *[]Like
	Count int
}

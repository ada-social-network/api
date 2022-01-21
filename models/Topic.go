package models

import uuid "github.com/satori/go.uuid"

// Topic model
type Topic struct {
	Base
	Name       string    `json:"name" binding:"required"`
	Content    string    `json:"content" binding:"required,min=4,max=21474"`
	UserID     uuid.UUID `gorm:"type=uuid" json:"userId"`
	CategoryID uuid.UUID `gorm:"type=uuid" json:"categoryId"`
	Posts      []Post    `json:"posts"`
}

package models

import uuid "github.com/satori/go.uuid"

// Topic model
type Topic struct {
	Base
	Name       string    `json:"name" binding:"required"`
	UserID     uuid.UUID `gorm:"type=uuid" json:"userId"`
	CategoryID uuid.UUID `gorm:"type=uuid" json:"categoryId"`
	Posts      []Post    `json:"posts"`
}

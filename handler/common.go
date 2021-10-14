package handler

import (
	"time"

	"gorm.io/gorm"
)

// CommonResource used for Gorm resources
type CommonResource struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

package models

import uuid "github.com/satori/go.uuid"

// Promo define a promo resource
type Promo struct {
	Base
	PromoName string `json:"promo"`
	StartDate string `json:"dateOfStart"`
	EndDate   string `json:"dateOfEnd"`
	Bio       string `json:"biography"`
	// By default, gorm will try to use UserID as a foreign key to the model User
	UserID uuid.UUID `gorm:"type=uuid" json:"userId"`
}

package models

import "gorm.io/gorm"

// Promo model dnejfbhjr
type Promo struct {
	gorm.Model
	PromoName string `json:"promo_name"`
	StartDate string `json:"date_of_start"`
	EndDate   string `json:"date_of_end"`
	Bio       string `json:"biography"`
}

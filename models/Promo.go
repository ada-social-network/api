package models

// Promo define a promo resource
type Promo struct {
	Base
	PromoName string `json:"promo"`
	StartDate string `json:"dateOfStart"`
	EndDate   string `json:"dateOfEnd"`
	Bio       string `json:"biography"`
}

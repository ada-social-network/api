package models

// Promo define a promo resource
type Promo struct {
	Base
	PromoName string `json:"promo_name"`
	StartDate string `json:"date_of_start"`
	EndDate   string `json:"date_of_end"`
	Bio       string `json:"biography"`
}

package models

// Promo define a promo resource
type Promo struct {
	Base
	Name       string `json:"name"`
	StartDate  string `json:"dateOfStart"`
	EndDate    string `json:"dateOfEnd"`
	Bio        string `json:"biography"`
	ProfilePic string `json:"profilePic"`
	// By default, gorm will try to use UserID as a foreign key to the model User
	Users []User `json:"users"`
}

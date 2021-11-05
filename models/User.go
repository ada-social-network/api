package models

import "gorm.io/gorm"

// User define a user resource
type User struct {
	gorm.Model
	LastName    string `json:"last_name" binding:"required,min=2,max=20"`
	FirstName   string `json:"first_name" binding:"required,min=2,max=20"`
	Email       string `json:"email" binding:"required,email"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

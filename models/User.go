package models

import "gorm.io/gorm"

// User define a user resource
type User struct {
	gorm.Model
	LastName       string    `json:"last_name" binding:"required,min=2,max=20"`
	FirstName      string    `json:"first_name" binding:"required,min=2,max=20"`
	Email          string    `json:"email" binding:"required,email"`
	DateOfBirth    string    `json:"date_of_birth" binding:"required"`
	Apprenticeship string    `json:"apprentice_at"`
	ProfilPic      string    `json:"profil_pic"`
	PrivateMail    string    `json:"private_mail"`
	Instagram      string    `json:"instagram"`
	Facebook       string    `json:"facebook"`
	Github         string    `json:"github"`
	Linkedin       string    `json:"linkedin"`
	MBTI           string    `json:"mbti"`
	Admin          bool      `json:"is_admin"`
	PromoID        uint      `json:"promo_id"`
	BdaPost        []BdaPost `gorm:"foreignKey:UserID"`
	Post           []Post    `gorm:"foreignKey:UserID"`
}

package models

import (
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// User define a user resource
type User struct {
	gorm.Model
	LastName       string    `json:"last_name" binding:"required,min=2,max=20"`
	FirstName      string    `json:"first_name" binding:"required,min=2,max=20"`
	Email          string    `json:"email" binding:"required,email" gorm:"unique"`
	Password       string    `json:"-"`
	DateOfBirth    string    `json:"date_of_birth"`
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
	BdaPosts       []BdaPost `json:"bda_posts"`
	Posts          []Post    `json:"posts"`
	Comments       []Comment `json:"comments"`
}

//HashPassword substitutes User.Password with its bcrypt hash
func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

//ComparePassword compares User.Password hash with raw password
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

//BeforeSave gorm hook
func (user *User) BeforeSave(db *gorm.DB) (err error) {
	if user.Password != "" {
		return user.HashPassword()
	}
	return nil
}

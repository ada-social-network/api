package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User define a user resource
type User struct {
	Base
	LastName       string    `json:"lastName" binding:"required,min=2,max=20"`
	FirstName      string    `json:"firstName" binding:"required,min=2,max=20"`
	Email          string    `json:"email" binding:"required,email" gorm:"unique"`
	Password       string    `json:"password"`
	DateOfBirth    string    `json:"dateOfBirth"`
	Apprenticeship string    `json:"apprenticeAt"`
	ProfilPic      string    `json:"profilPic"`
	Biography      string    `json:"biography"`
	CoverPic       string    `json:"coverPic"`
	PrivateMail    string    `json:"privateMail"`
	ProjectPerso   string    `json:"projectPerso"`
	ProjectPro     string    `json:"projectPro"`
	Instagram      string    `json:"instagram"`
	Facebook       string    `json:"facebook"`
	Github         string    `json:"github"`
	Linkedin       string    `json:"linkedin"`
	MBTI           string    `json:"mbti"`
	Admin          bool      `json:"isAdmin"`
	PromoID        uuid.UUID `gorm:"type=uuid" json:"promoId"`
	BdaPosts       []BdaPost `json:"bdaPosts"`
	Posts          []Post    `json:"posts"`
	Comments       []Comment `json:"comments"`
	Topics         []Topic   `json:"topics"`
	Likes          []Like    `json:"likes"`
}

//ComparePassword compares User.Password hash with raw password
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// HashPassword hashes password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

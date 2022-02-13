package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserResponse define a user response
type UserResponse struct {
	models.Base
	LastName       string           `json:"lastName" binding:"required,min=2,max=20"`
	FirstName      string           `json:"firstName" binding:"required,min=2,max=20"`
	Email          string           `json:"email" binding:"required,email" gorm:"unique"`
	DateOfBirth    string           `json:"dateOfBirth"`
	Apprenticeship string           `json:"apprenticeAt"`
	ProfilPic      string           `json:"profilPic"`
	Biography      string           `json:"biography"`
	CoverPic       string           `json:"coverPic"`
	PrivateMail    string           `json:"privateMail"`
	ProjectPerso   string           `json:"projectPerso"`
	ProjectPro     string           `json:"projectro"`
	Instagram      string           `json:"instagram"`
	Facebook       string           `json:"facebook"`
	Github         string           `json:"github"`
	Linkedin       string           `json:"linkedin"`
	MBTI           string           `json:"mbti"`
	Admin          bool             `json:"isAdmin"`
	PromoID        uuid.UUID        `gorm:"type=uuid" json:"promoId"`
	BdaPosts       []models.BdaPost `json:"bdaPosts"`
	Posts          []models.Post    `json:"posts"`
	Comments       []models.Comment `json:"comments"`
	Topics         []models.Topic   `json:"topics"`
	Likes          []models.Like    `json:"likes"`
}

// ListUser respond a list of users
func ListUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := &[]models.User{}

		result := db.Find(&users)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, createUsersResponse(users))
	}
}

// CreateUser create a user
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{}
		err := c.ShouldBindJSON(user)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		user.Password, err = models.HashPassword(user.Password)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(user)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, user)
	}
}

// DeleteUser delete a specific user
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.User{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetUser get a specific user
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		user := &models.User{}

		result := db.First(user, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "User", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, createUserResponse(user))
	}
}

// UpdateUser update a specific user
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		user := &models.User{}

		result := db.Omit("Password").First(user, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "User", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(user)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		// we omit password because if a hashed password is present it will be re-encrypted
		if user.Password == "" {
			db.Omit("Password").Save(user)
		} else {
			hashedPassword, err := models.HashPassword(user.Password)
			if err != nil {
				httpError.Internal(c, err)
				return
			}

			user.Password = hashedPassword
			db.Save(user)
		}

		c.JSON(200, createUserResponse(user))
	}
}

// createUserResponse map the values of user to createUserResponse
func createUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Base:           user.Base,
		LastName:       user.LastName,
		FirstName:      user.FirstName,
		Email:          user.Email,
		DateOfBirth:    user.DateOfBirth,
		Apprenticeship: user.Apprenticeship,
		ProfilPic:      user.ProfilPic,
		Biography:      user.Biography,
		CoverPic:       user.CoverPic,
		PrivateMail:    user.PrivateMail,
		ProjectPerso:   user.ProjectPerso,
		ProjectPro:     user.ProjectPro,
		Instagram:      user.Instagram,
		Facebook:       user.Facebook,
		Github:         user.Github,
		Linkedin:       user.Linkedin,
		MBTI:           user.MBTI,
		Admin:          user.Admin,
		PromoID:        user.PromoID,
		BdaPosts:       user.BdaPosts,
		Posts:          user.Posts,
		Comments:       user.Comments,
		Topics:         user.Topics,
		Likes:          user.Likes,
	}
}

// createUserResponse map the values of all users to a list of createUserResponse
func createUsersResponse(users *[]models.User) []UserResponse {
	usersList := []UserResponse{}
	for _, u := range *users {
		usersList = append(usersList, createUserResponse(&u))
	}

	return usersList
}

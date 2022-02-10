package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// MeHandler provide informations about the connected user
func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("id")
		if !exist {
			httpError.Internal(c, errors.New("jwt does not exist"))
			return
		}

		u, ok := user.(*models.User)
		if !ok {
			httpError.Internal(c, errors.New("this is not a user"))
			return
		}

		tx := db.First(u, "id = ?", u.ID)
		if tx.Error != nil {
			httpError.Internal(c, tx.Error)
			return
		}

		c.JSON(200, u)
	}
}

// UpdatePasswordRequest is the request for the password change
type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

// UpdatePassword update a specific user
func UpdatePassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
		}

		updatePasswordRequest := &UpdatePasswordRequest{}
		err = c.ShouldBindJSON(updatePasswordRequest)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		hashedPassword, err := models.HashPassword(updatePasswordRequest.Password)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		tx := db.Model(&user).Update("password", hashedPassword)
		if tx.Error != nil {
			httpError.Internal(c, err)
		}

		c.JSON(204, nil)
	}
}

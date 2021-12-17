package handler

import (
	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userRegister struct {
	LastName  string `json:"lastName" binding:"required,min=2,max=20"`
	FirstName string `json:"firstName" binding:"required,min=2,max=20"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=32"`
}

// Register register a user
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRegister := &userRegister{}

		err := c.BindJSON(userRegister)
		if err != nil {
			ve, ok := err.(validator.ValidationErrors)
			if ok {
				httpError.Validation(c, ve)
				return
			}

			httpError.Internal(c, err)
			return
		}

		user := &models.User{
			LastName:  userRegister.LastName,
			FirstName: userRegister.FirstName,
			Email:     userRegister.Email,
			Password:  userRegister.Password,
		}

		tx := db.First(&models.User{}, "email = ?", user.Email)
		if tx.RowsAffected != 0 {
			httpError.AlreadyExist(c, "email", user.Email)
			return
		}

		tx = db.Create(user)
		if tx.Error != nil {
			httpError.Internal(c, tx.Error)
			return
		}

		c.JSON(200, user)
	}
}

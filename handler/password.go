package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateUserPassword update a specific user
func UpdateUserPassword(db *gorm.DB) gin.HandlerFunc {
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
			db.Save(user)
		}

		c.JSON(200, user)
	}
}

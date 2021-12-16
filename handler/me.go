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

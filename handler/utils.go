package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/ada-social-network/api/middleware"
	"github.com/ada-social-network/api/models"
)

// GetCurrentUser Get the current connected user
func GetCurrentUser(c *gin.Context) (*models.User, error) {
	user, _ := c.Get(middleware.IdentityKey)
	u, ok := user.(*models.User)
	if !ok {
		return nil, errors.New("not a user")
	}

	if u == nil {
		return nil, errors.New("missing user")
	}

	return u, nil
}

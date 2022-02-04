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

// Collection defines the count of items and items
type Collection struct {
	Items            []interface{} `json:"items"`
	Count            int           `json:"count"`
	isLikedByCurrent bool          `json:"isLiked"`
}

// NewCollection create a new collection
func NewCollection(items []interface{}) *Collection {
	return &Collection{Items: items, Count: len(items)}
}

func NewLikeCollection(items []interface{}, isLikedByCurrent bool) *Collection {
	return &Collection{Items: items, Count: len(items), isLikedByCurrent: isLikedByCurrent}
}

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
	Items []interface{} `json:"items"`
	Count int           `json:"count"`
}

//LikeCollection defines the count of items and items and if is liked by current user
type LikeCollection struct {
	Items                []interface{} `json:"items"`
	Count                int           `json:"count"`
	IsLikedByCurrentUser bool          `json:"isLikedByCurrentUser"`
}

// NewCollection create a new collection
func NewCollection(items []interface{}) *Collection {
	return &Collection{Items: items, Count: len(items)}
}

// NewLikeCollection create a new collection with a boolean for likes
func NewLikeCollection(items []interface{}, isLikedByCurrentUser bool) *LikeCollection {
	return &LikeCollection{Items: items, Count: len(items), IsLikedByCurrentUser: isLikedByCurrentUser}
}

package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Post define a post resource
type Post struct {
	CommonResource
	Content string `json:"content" binding:"required,min=4,max=1024"`
}

// ListPostHandler respond a list of posts
func ListPostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := &[]Post{}

		result := db.Find(&posts)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, posts)
	}
}

// CreatePostHandler create a post
func CreatePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post := &Post{}

		err := c.ShouldBindJSON(post)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(post)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, post)
	}
}

// DeletePostHandler delete a specific post
func DeletePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&Post{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetPostHandler get a specific post
func GetPostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		post := &Post{}

		result := db.First(post, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Post", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, post)
	}
}

// UpdatePostHandler update a specific post
func UpdatePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		post := &Post{}

		result := db.First(post, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Post", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(post)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(post)

		c.JSON(200, post)
	}
}

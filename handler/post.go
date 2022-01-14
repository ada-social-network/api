package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ListPostHandler respond a list of posts
func ListPostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := &[]models.Post{}

		result := db.Find(&posts)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, posts)
	}
}

// CreatePost create a post
func CreatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		id, _ := c.Params.Get("id")
		post := &models.Post{}

		err = c.ShouldBindJSON(post)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		post.UserID = user.ID
		topicUuid, err := uuid.FromString(id)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		post.TopicID = topicUuid

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

		result := db.Delete(&models.Post{}, "id = ?", id)
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
		post := &models.Post{}

		result := db.First(post, "id = ?", id)
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
		post := &models.Post{}

		result := db.First(post, "id = ?", id)
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

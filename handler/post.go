package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

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
		topicUUID, err := uuid.FromString(id)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		post.TopicID = topicUUID

		result := db.Create(post)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, post)
	}
}

// DeletePost delete a specific post
func DeletePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, _ := c.Params.Get("postId")

		result := db.Delete(&models.Post{}, "id = ?", postID)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetPost get a specific post
func GetPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, _ := c.Params.Get("postID")
		post := &models.Post{}

		result := db.First(post, "id = ?", postID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Post", postID, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, post)
	}
}

// UpdatePost update a specific post
func UpdatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, _ := c.Params.Get("postID")
		post := &models.Post{}

		result := db.First(post, "id = ?", postID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Post", postID, result.Error)
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

// ListPosts get posts of a topic
func ListPosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		posts := &[]models.Post{}

		result := db.Find(posts, "topic_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, posts)
	}
}

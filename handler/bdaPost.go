package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListbdaPost respond a list of posts
func ListBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bdaPosts := &[]models.BdaPost{}

		result := db.Find(&bdaPosts)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, bdaPosts)
	}
}

// CreatePostHandler create a post
func CreateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bdaPost := &models.BdaPost{}

		err := c.ShouldBindJSON(bdaPost)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(bdaPost)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, bdaPost)
	}
}

// DeletePostHandler delete a specific post
func DeleteBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.BdaPost{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetPostHandler get a specific post
func GetBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}

		result := db.First(bdaPost, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "BdaPost", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, bdaPost)
	}
}

// UpdatePostHandler update a specific post
func UpdateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}

		result := db.First(bdaPost, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "BdaPost", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(bdaPost)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(bdaPost)

		c.JSON(200, bdaPost)
	}
}

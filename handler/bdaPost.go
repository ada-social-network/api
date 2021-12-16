package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// ListBdaPost respond a list of bda posts
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

// CreateBdaPost create a bda post
func CreateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPost := &models.BdaPost{}

		err = c.ShouldBindJSON(bdaPost)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPost.UserID = user.ID

		result := db.Create(bdaPost)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, bdaPost)
	}
}

// DeleteBdaPost delete a specific bda post
func DeleteBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.BdaPost{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetBdaPost get a specific bda post
func GetBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}

		result := db.First(bdaPost, "id = ?", id)
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

// UpdateBdaPost update a specific bda post
func UpdateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}

		result := db.First(bdaPost, "id = ?", id)
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

// ListBdaPostComments get comments of a bda post
func ListBdaPostComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		comments := &[]models.Comment{}

		result := db.Find(comments, "bdapost_id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, comments)
	}
}

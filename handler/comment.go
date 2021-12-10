package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// ListComment respond a list of comments
func ListComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		comments := &[]models.Comment{}

		result := db.Find(&comments)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, comments)
	}
}

// CreateComment create a comment
func CreateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		comments := &models.Comment{}
		err := c.ShouldBindJSON(comments)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(comments)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, comments)
	}
}

// DeleteComment delete a specific comment
func DeleteComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		result := db.Delete(&models.Comment{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetComment get a specific comment
func GetComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		comments := &models.Comment{}

		result := db.First(comments, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Comment", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, comments)
	}
}

// UpdateComment update a specific comment
func UpdateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		comments := &models.Comment{}

		result := db.First(comments, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Comment", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(comments)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(comments)

		c.JSON(200, comments)
	}
}

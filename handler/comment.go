package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListComment respond a list of users
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

// CreateComment create a user
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

// DeleteComment delete a specific user
func DeleteComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		result := db.Delete(&models.Comment{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetComment get a specific user
func GetComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		comments := &models.Comment{}

		result := db.First(comments, id)
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

// UpdateComment update a specific user
func UpdateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		comments := &models.Comment{}

		result := db.First(comments, id)
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

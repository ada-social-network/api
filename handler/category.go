package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// ListCategories respond a list of categories
func ListCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := &[]models.Category{}

		result := db.Find(&categories)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, categories)
	}
}

// CreateCategory create a category
func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		category := &models.Category{}

		err := c.ShouldBindJSON(category)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(category)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, category)
	}
}

// DeleteCategory delete a specific category
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.Category{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// ListCategoryTopics get topics of a category
func ListCategoryTopics(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		topics := &[]models.Topic{}

		result := db.Find(topics, "category_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, topics)
	}
}

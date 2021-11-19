package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListPromo respond a list of promo
func ListPromo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		promos := &[]models.Promo{}

		result := db.Find(&promos)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, promos)
	}
}

// CreatePromo create a promo
func CreatePromo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		promo := &models.Promo{}

		err := c.ShouldBindJSON(promo)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(promo)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, promo)
	}
}

// DeletePromo delete a specific promo
func DeletePromo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.Promo{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// UpdatePromo update a specific promo
func UpdatePromo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		promo := &models.Promo{}

		result := db.First(promo, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Promo", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(promo)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(promo)

		c.JSON(200, promo)
	}
}

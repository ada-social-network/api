package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// ListPromoUsers get users of a promo
func ListPromoUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		promoUsers := &[]models.Promo{}

		result := db.Find(promoUsers, "promo_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, promoUsers)
	}
}

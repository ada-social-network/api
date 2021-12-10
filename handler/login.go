package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/ada-social-network/api/models"
)

// MeHandler provide informations about the connected user
func MeHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")
	c.JSON(200, gin.H{
		"userID":    claims["ID"],
		"admin":     claims["admin"],
		"userEmail": user.(*models.User).Email,
	})
}

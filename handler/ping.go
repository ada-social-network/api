package handler

import "github.com/gin-gonic/gin"

// PingHandler respond dummy message
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

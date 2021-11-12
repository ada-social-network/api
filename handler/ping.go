package handler

import "github.com/gin-gonic/gin"

// Pong is a dummy response
type Pong struct {
	Message string `json:"message"`
}

// Ping respond dummy message
func Ping(c *gin.Context) {
	c.JSON(200, Pong{Message: "pong"})
}

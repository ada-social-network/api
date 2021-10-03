package handler

import "github.com/gin-gonic/gin"

// Pong is a dummy response
type Pong struct {
	Message string `json:"message"`
}

// PingHandler respond dummy message
func PingHandler(c *gin.Context) {
	c.JSON(200, Pong{Message: "pong"})
}

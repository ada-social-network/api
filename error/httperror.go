package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Content of an error
type Content struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message"`
}

// Internal respond with an internal error
func Internal(c *gin.Context, err error) {
	HTTPError(c, http.StatusInternalServerError, "Internal error", err)
}

// HTTPError respond an http generic error
func HTTPError(c *gin.Context, code int, msg string, err error) {
	mode := gin.Mode()
	content := Content{
		Message: msg,
	}

	if mode == gin.DebugMode {
		content.Error = err.Error()
	}

	c.JSON(code, content)
}

package error

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Content of an error
type Content struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message"`
}

// AlreadyExist respond with an already exist error
func AlreadyExist(c *gin.Context, field, value string) {
	Conflict(c, errors.New("this email is already taken"))
}

// AlreadyLiked respond with an already liked error
func AlreadyLiked(c *gin.Context, field, value string) {
	Conflict(c, errors.New("resource already liked by this user"))
}

// Conflict respond with a conflict error
func Conflict(c *gin.Context, err error) {
	HTTPError(c, http.StatusConflict, err.Error(), err)
}

// Internal respond with an internal error
func Internal(c *gin.Context, err error) {
	HTTPError(c, http.StatusInternalServerError, "Internal error", err)
}

// Validation respond with a validation error
func Validation(c *gin.Context, err validator.ValidationErrors) {
	HTTPError(c, http.StatusBadRequest, err.Error(), err)
}

// NotFound respond with a not found error
func NotFound(c *gin.Context, name string, id string, err error) {
	HTTPError(
		c,
		http.StatusNotFound,
		fmt.Sprintf("Resource %s with id %s can not found", name, id),
		err,
	)
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

package middleware

import "github.com/gin-gonic/gin"

const (
	versionHeader = "X-version"
)

// Version Set version in header
func Version(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(versionHeader, version)
	}
}

package handler

import (
	_ "embed"
	"encoding/json"

	httpError "github.com/ada-social-network/api/error"
	"github.com/gin-gonic/gin"
)

var (
	//go:embed fixtures/list-posts.json
	listPosts []byte
)

// Post define a post resource
type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// PostHandler respond a list of posts
func PostHandler(c *gin.Context) {
	posts := &[]Post{}

	err := json.Unmarshal(listPosts, posts)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, posts)
}

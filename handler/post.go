package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	httpError "github.com/ada-social-network/api/error"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Post define a post resource
type Post struct {
	CommonResource
	Content string `json:"content"`
}

// ListHandler respond a list of posts
func ListHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := &[]Post{}

		result := db.Find(&posts)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, posts)
	}
}

// CreateHandler create a post
func CreateHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		post := &Post{}
		err = json.Unmarshal(jsonData, post)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(post)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, post)
	}
}

// DeleteHandler delete a specific post
func DeleteHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id := c.Query("id")
		result := db.Delete(&Post{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetPostHandler get a specific post
func GetPostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		post := &Post{}

		result := db.First(post, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Post", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, post)
	}
}

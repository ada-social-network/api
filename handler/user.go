package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	httpError "github.com/ada-social-network/api/error"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User define a user resource
type User struct {
	CommonResource
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}

// ListUserHandler respond a list of users
func ListUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := &[]User{}

		result := db.Find(&users)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, users)
	}
}

// CreateUserHandler create a user
func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		user := &User{}
		err = json.Unmarshal(jsonData, user)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Create(user)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, user)
	}
}

// DeleteUserHandler delete a specific user
func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id := c.Query("id")
		result := db.Delete(&User{}, id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetUserHandler get a specific user
func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		id, _ := c.Params.Get("id")
		user := &User{}

		result := db.First(user, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "User", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, user)
	}
}

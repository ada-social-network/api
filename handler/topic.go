package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
)

// GetTopic get a specific bda post
func GetTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		topic := &models.Topic{}

		result := db.First(topic, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Topic", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, topic)
	}
}

// ListTopics respond a list of topics
func ListTopics(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		topics := &[]models.Topic{}

		result := db.Find(&topics)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, topics)
	}
}

// CreateTopic create a topic
func CreateTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		topic := &models.Topic{}
		id, _ := c.Params.Get("id")

		err = c.ShouldBindJSON(topic)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		topic.UserID = user.ID
		catUUID, err := uuid.FromString(id)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		topic.CategoryID = catUUID

		result := db.Create(topic)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, topic)
	}
}

// DeleteTopic delete a specific topic and maybe we need to implement that it deletes all the posts of this topics ?
func DeleteTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.Topic{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// UpdateTopic update a specific bda post
func UpdateTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		topic := &models.Topic{}

		result := db.First(topic, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Topic", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(topic)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(topic)

		c.JSON(200, topic)
	}
}

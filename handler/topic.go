package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// TopicHandler is a struct to define bda post handler
type TopicHandler struct {
	repository *repository.TopicRepository
}

// NewTopicHandler is a factory for topic handler
func NewTopicHandler(repository *repository.TopicRepository) *TopicHandler {
	return &TopicHandler{repository: repository}
}

// GetTopic get a specific topic
func (t *TopicHandler) GetTopic(c *gin.Context) {
	id, _ := c.Params.Get("id")

	topic := &models.Topic{}

	err := t.repository.GetTopicByID(topic, id)
	if err != nil {
		if errors.Is(err, repository.ErrTopicNotFound) {
			httpError.NotFound(c, "topic", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, topic)
}

// ListTopics respond a list of topics
func (t *TopicHandler) ListTopics(c *gin.Context) {
	topics := &[]models.Topic{}

	err := t.repository.ListAllTopics(topics)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, topics)
}

// CreateTopic create a topic
func (t *TopicHandler) CreateTopic(c *gin.Context) {
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	topic := &models.Topic{}
	categoryID, _ := c.Params.Get("id")

	err = c.ShouldBindJSON(topic)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	topic.UserID = user.ID
	catUUID, err := uuid.FromString(categoryID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}
	topic.CategoryID = catUUID

	err = t.repository.CreateTopic(topic)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, topic)
}

// DeleteTopic delete a specific topic
func (t *TopicHandler) DeleteTopic(c *gin.Context) {
	id, _ := c.Params.Get("id")

	err := t.repository.DeleteTopicByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrTopicNotFound) {
			httpError.NotFound(c, "topic", id, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

// UpdateTopic update a specific topic
func (t *TopicHandler) UpdateTopic(c *gin.Context) {
	id, _ := c.Params.Get("id")
	topic := &models.Topic{}

	err := t.repository.GetTopicByID(topic, id)
	if err != nil {
		if errors.Is(err, repository.ErrTopicNotFound) {
			httpError.NotFound(c, "topic", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(topic)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = t.repository.UpdateTopic(topic)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, topic)
}

// ListCategoryTopics get topics of a category
func (t *TopicHandler) ListCategoryTopics(c *gin.Context) {
	categoryID, _ := c.Params.Get("id")
	topics := &[]models.Topic{}

	err := t.repository.ListAllTopicsByCategoryID(topics, categoryID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, topics)
}

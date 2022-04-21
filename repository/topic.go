package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	"gorm.io/gorm"
)

// ErrTopicNotFound is an error when resource is not found
var (
	ErrTopicNotFound = errors.New("topic not found")
)

// TopicRepository is a repository for topic resource
type TopicRepository struct {
	db *gorm.DB
}

// NewTopicRepository is to create a new topic repository
func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

// CreateTopic create a topic in the DB
func (t *TopicRepository) CreateTopic(topic *models.Topic) error {
	return t.db.Create(topic).Error
}

// ListAllTopics list all the topics in the DB
func (t *TopicRepository) ListAllTopics(topics *[]models.Topic) error {
	return t.db.Find(topics).Error
}

// ListAllTopicsByCategoryID list all the topics from a specific category in the DB
func (t *TopicRepository) ListAllTopicsByCategoryID(topics *[]models.Topic, categoryID string) error {
	return t.db.Find(topics, "category_id= ?", categoryID).Error
}

// GetTopicByID get a bda post by id in the DB
func (t *TopicRepository) GetTopicByID(topic *models.Topic, topicID string) error {
	tx := t.db.First(topic, "id = ?", topicID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrTopicNotFound
	}

	return tx.Error
}

// UpdateTopic update a topic in the DB
func (t *TopicRepository) UpdateTopic(topic *models.Topic) error {
	return t.db.Save(topic).Error
}

// DeleteTopicByID delete a topic by ID in the DB
func (t *TopicRepository) DeleteTopicByID(topicID string) error {
	tx := t.db.Delete(&models.Topic{}, "id = ?", topicID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrTopicNotFound
	}

	return tx.Error
}

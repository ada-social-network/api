package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	"gorm.io/gorm"
)

// ErrPostNotFound is an error when resource is not found
var (
	ErrPostNotFound = errors.New("post not found")
)

// PostRepository is a repository for post resource
type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository is to create a new post repository
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// CreatePost create a post in the DB
func (p *PostRepository) CreatePost(Post *models.Post) error {
	return p.db.Create(Post).Error
}

// ListAllPosts list all the post in the DB
func (p *PostRepository) ListAllPosts(posts *[]models.Post) error {
	return p.db.Find(posts).Error
}

// ListAllPostsByTopicID list all posts of a specific topic in the DB
func (p *PostRepository) ListAllPostsByTopicID(posts *[]models.Post, topicID string) error {
	return p.db.Find(posts, "topic_id=?", topicID).Error
}

// GetPostByID get a post by id in the DB
func (p *PostRepository) GetPostByID(post *models.Post, postID string) error {
	tx := p.db.First(post, "id = ?", postID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrPostNotFound
	}

	return tx.Error
}

// UpdatePost update a post in the DB
func (p *PostRepository) UpdatePost(post *models.Post) error {
	return p.db.Save(post).Error
}

// DeletePostByID delete a post by ID in the DB
func (p *PostRepository) DeletePostByID(postID string) error {
	tx := p.db.Delete(&models.Post{}, "id = ?", postID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrPostNotFound
	}

	return tx.Error
}

package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	uuid "github.com/satori/go.uuid"
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

// CheckLikeByUserAndPostID will check if a like by this user already exist on a post
func (p *PostRepository) CheckLikeByUserAndPostID(like *models.Like, userID uuid.UUID, postID uuid.UUID) (bool, error) {
	tx := p.db.Where("user_id= ? AND post_id= ?", userID, postID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

// CreateLike create a like on a resource in the DB
func (p *PostRepository) CreateLike(like *models.Like) error {
	return p.db.Create(like).Error
}

// ListAllPostsByPostID list all likes of a specific post in the DB
func (p *PostRepository) ListAllPostsByPostID(likes *[]models.Like, postID string) error {
	return p.db.Find(likes, "post_id=?", postID).Error
}

// DeleteLikeByID delete a like by ID in the DB
func (p *PostRepository) DeleteLikeByID(likeID string) error {
	tx := p.db.Delete(&models.Like{}, "id = ?", likeID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrLikeNotFound
	}

	return tx.Error
}

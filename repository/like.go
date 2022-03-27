package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ErrLikeNotFound is an error when resource is not found
var (
	ErrLikeNotFound = errors.New("like not found")
)

// LikeRepository is a repository for like resource
type LikeRepository struct {
	db *gorm.DB
}

// NewLikeRepository is to create a new like repository
func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

// CreateLike create a like on a resource in the DB
func (l *LikeRepository) CreateLike(like *models.Like) error {
	return l.db.Create(like).Error
}

// ListAllPostsByPostID list all likes of a specific post in the DB
func (l *LikeRepository) ListAllPostsByPostID(likes *[]models.Like, postID string) error {
	return l.db.Find(likes, "post_id=?", postID).Error
}

// DeleteLikeByID delete a like by ID in the DB
func (l *LikeRepository) DeleteLikeByID(likeID string) error {
	tx := l.db.Delete(&models.Like{}, "id = ?", likeID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrLikeNotFound
	}

	return tx.Error
}

// CheckLikeByUserAndPostID will check if a like by this user already exist on a resource
func (l *LikeRepository) CheckLikeByUserAndPostID(like *models.Like, userID uuid.UUID, postID uuid.UUID) (bool, error) {
	tx := l.db.Where("user_id= ? AND post_id= ?", userID, postID).Find(like)
	if tx.RowsAffected > 0 {
		return true, nil
	}
	return false, tx.Error
}

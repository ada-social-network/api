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

// ListAllPostsByBdaPostID list all likes of a specific bda post in the DB
func (l *LikeRepository) ListAllPostsByBdaPostID(likes *[]models.Like, bdaPostID string) error {
	return l.db.Find(likes, "bda_post_id=?", bdaPostID).Error
}

// ListAllPostsByCommentID list all likes of a specific comment in the DB
func (l *LikeRepository) ListAllPostsByCommentID(likes *[]models.Like, commentID string) error {
	return l.db.Find(likes, "comment_id=?", commentID).Error
}

// DeleteLikeByID delete a like by ID in the DB
func (l *LikeRepository) DeleteLikeByID(likeID string) error {
	tx := l.db.Delete(&models.Like{}, "id = ?", likeID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrLikeNotFound
	}

	return tx.Error
}

// CheckLikeByUserAndPostID will check if a like by this user already exist on a post
func (l *LikeRepository) CheckLikeByUserAndPostID(like *models.Like, userID uuid.UUID, postID uuid.UUID) (bool, error) {
	tx := l.db.Where("user_id= ? AND post_id= ?", userID, postID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

// CheckLikeByUserAndBdaPostID will check if a like by this user already exist on a bda post
func (l *LikeRepository) CheckLikeByUserAndBdaPostID(like *models.Like, userID uuid.UUID, bdaPostID uuid.UUID) (bool, error) {
	tx := l.db.Where("user_id= ? AND bda_post_id= ?", userID, bdaPostID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

// CheckLikeByUserAndCommentID will check if a like by this user already exist on a comment
func (l *LikeRepository) CheckLikeByUserAndCommentID(like *models.Like, userID uuid.UUID, commentID uuid.UUID) (bool, error) {
	tx := l.db.Where("user_id= ? AND comment_id= ?", userID, commentID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

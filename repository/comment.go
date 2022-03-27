package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ErrCommentNotFound is an error when resource is not found
var (
	ErrCommentNotFound = errors.New("comment not found")
)

// CommentRepository is a repository for comment resource
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository is to create a new comment repository
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// GetCommentByID get a comment by id in the DB
func (co *CommentRepository) GetCommentByID(comment *models.Comment, commentID string) error {
	tx := co.db.First(comment, "id = ?", commentID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrCommentNotFound
	}

	return tx.Error
}

// ListAllComments list all comments in the DB
func (co *CommentRepository) ListAllComments(comments *[]models.Comment) error {
	return co.db.Find(comments).Error
}

// ListAllCommentsByBdaPostID list all comments of a specific bdaPost in the DB
func (co *CommentRepository) ListAllCommentsByBdaPostID(comments *[]models.Comment, bdaPostID string) error {
	return co.db.Find(comments, "bda_post_id=?", bdaPostID).Error
}

// CreateComment create a comment in the DB
func (co *CommentRepository) CreateComment(comment *models.Comment) error {
	return co.db.Create(comment).Error
}

// UpdateComment update a comment in the DB
func (co *CommentRepository) UpdateComment(comment *models.Comment) error {
	return co.db.Save(comment).Error
}

// DeleteCommentByID delete a comment by ID in the DB
func (co *CommentRepository) DeleteCommentByID(commentID string) error {
	tx := co.db.Delete(&models.Comment{}, "id = ?", commentID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrCommentNotFound
	}

	return tx.Error
}

// CheckLikeByUserAndCommentID will check if a like by this user already exist on a comment
func (co *CommentRepository) CheckLikeByUserAndCommentID(like *models.Like, userID uuid.UUID, commentID uuid.UUID) (bool, error) {
	tx := co.db.Where("user_id= ? AND comment_id= ?", userID, commentID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

// CreateLike create a like on a resource in the DB
func (co *CommentRepository) CreateLike(like *models.Like) error {
	return co.db.Create(like).Error
}

// ListAllPostsByCommentID list all likes of a specific comment in the DB
func (co *CommentRepository) ListAllPostsByCommentID(likes *[]models.Like, commentID string) error {
	return co.db.Find(likes, "comment_id=?", commentID).Error
}

// DeleteLikeByID delete a like by ID in the DB
func (co *CommentRepository) DeleteLikeByID(likeID string) error {
	tx := co.db.Delete(&models.Like{}, "id = ?", likeID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrLikeNotFound
	}

	return tx.Error
}

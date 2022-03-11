package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
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

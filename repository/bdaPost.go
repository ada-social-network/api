package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ErrBdaPostNotFound is an error when resource is not found
var (
	ErrBdaPostNotFound = errors.New("bda post not found")
)

// BdaPostRepository is a repository for bda post resource
type BdaPostRepository struct {
	db *gorm.DB
}

// NewBdaPostRepository is to create a new bda post repository
func NewBdaPostRepository(db *gorm.DB) *BdaPostRepository {
	return &BdaPostRepository{db: db}
}

// CreateBdaPost create a bda post in the DB
func (bp *BdaPostRepository) CreateBdaPost(bdaPost *models.BdaPost) error {
	return bp.db.Create(bdaPost).Error
}

// ListAllBdaPosts list all the bda post in the DB
func (bp *BdaPostRepository) ListAllBdaPosts(bdaPosts *[]models.BdaPost) error {
	return bp.db.Find(bdaPosts).Error
}

// GetBdaPostByID get a bda post by id in the DB
func (bp *BdaPostRepository) GetBdaPostByID(bdaPost *models.BdaPost, bdaPostID string) error {
	tx := bp.db.First(bdaPost, "id = ?", bdaPostID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrBdaPostNotFound
	}

	return tx.Error
}

// UpdateBdaPost update a bda post in the DB
func (bp *BdaPostRepository) UpdateBdaPost(bdaPost *models.BdaPost) error {
	return bp.db.Save(bdaPost).Error
}

// DeleteBdaPostByID delete a bda post by ID in the DB
func (bp *BdaPostRepository) DeleteBdaPostByID(bdaPostID string) error {
	tx := bp.db.Delete(&models.BdaPost{}, "id = ?", bdaPostID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrBdaPostNotFound
	}

	return tx.Error
}

// CheckLikeByUserAndBdaPostID will check if a like by this user already exist on a bda post
func (bp *BdaPostRepository) CheckLikeByUserAndBdaPostID(like *models.Like, userID uuid.UUID, bdaPostID uuid.UUID) (bool, error) {
	tx := bp.db.Where("user_id= ? AND bda_post_id= ?", userID, bdaPostID).First(like)
	return tx.RowsAffected > 0, tx.Error
}

// CreateLike create a like on a resource in the DB
func (bp *BdaPostRepository) CreateLike(like *models.Like) error {
	return bp.db.Create(like).Error
}

// ListAllPostsByBdaPostID list all likes of a specific bda post in the DB
func (bp *BdaPostRepository) ListAllPostsByBdaPostID(likes *[]models.Like, bdaPostID string) error {
	return bp.db.Find(likes, "bda_post_id=?", bdaPostID).Error
}

// DeleteLikeByID delete a like by ID in the DB
func (bp *BdaPostRepository) DeleteLikeByID(likeID string) error {
	tx := bp.db.Delete(&models.Like{}, "id = ?", likeID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrLikeNotFound
	}

	return tx.Error
}

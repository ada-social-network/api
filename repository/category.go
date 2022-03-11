package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	"gorm.io/gorm"
)

// ErrCategoryNotFound is an error when resource is not found
var (
	ErrCategoryNotFound = errors.New("category not found")
)

//CategoryRepository is a repository for category resource
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository is to create a new category repository
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// CreateCategory create a category in the DB
func (ca *CategoryRepository) CreateCategory(category *models.Category) error {
	return ca.db.Create(category).Error
}

// ListAllCategories list all categories in the DB
func (ca *CategoryRepository) ListAllCategories(categories *[]models.Category) error {
	return ca.db.Find(categories).Error
}

// GetCategoryByID get a category by id in the DB
func (ca *CategoryRepository) GetCategoryByID(category *models.Category, categoryID string) error {
	tx := ca.db.First(category, "id = ?", categoryID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrCategoryNotFound
	}

	return tx.Error
}

// UpdateCategory update a category in the DB
func (ca *CategoryRepository) UpdateCategory(category *models.Category) error {
	return ca.db.Save(category).Error
}

// DeleteCategoryByID delete a category by ID in the DB
func (ca *CategoryRepository) DeleteCategoryByID(categoryID string) error {
	tx := ca.db.Delete(&models.Category{}, "id = ?", categoryID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrCategoryNotFound
	}

	return tx.Error
}

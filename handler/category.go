package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
)

//CategoryHandler is a struct to define category handler
type CategoryHandler struct {
	repository *repository.CategoryRepository
}

// NewCategoryHandler is a factory for category handler
func NewCategoryHandler(repository *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repository: repository}
}

// ListCategories respond a list of categories
func (ca *CategoryHandler) ListCategories(c *gin.Context) {
	categories := &[]models.Category{}

	err := ca.repository.ListAllCategories(categories)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, categories)
}

// CreateCategory create a category
func (ca *CategoryHandler) CreateCategory(c *gin.Context) {
	category := &models.Category{}

	err := c.ShouldBindJSON(category)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = ca.repository.CreateCategory(category)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, category)
}

// DeleteCategory delete a specific category
func (ca *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := c.Params.Get("id")

	err := ca.repository.DeleteCategoryByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			httpError.NotFound(c, "category", id, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

// GetCategory get a specific category
func (ca *CategoryHandler) GetCategory(c *gin.Context) {
	id, _ := c.Params.Get("id")

	category := &models.Category{}

	err := ca.repository.GetCategoryByID(category, id)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			httpError.NotFound(c, "category", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, category)
}

// UpdateCategory update a specific category
func (ca *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := c.Params.Get("id")
	category := &models.Category{}

	err := ca.repository.GetCategoryByID(category, id)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			httpError.NotFound(c, "category", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(category)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = ca.repository.UpdateCategory(category)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, category)
}

package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	"gorm.io/gorm"
)

// ErrPromoNotFound is an error when resource is not found
var (
	ErrPromoNotFound = errors.New("promo not found")
)

// PromoRepository is a repository for promo resource
type PromoRepository struct {
	db *gorm.DB
}

// NewPromoRepository is to create a new promo repository
func NewPromoRepository(db *gorm.DB) *PromoRepository {
	return &PromoRepository{db: db}
}

// GetPromoByID get a promo by id in the DB
func (p *PromoRepository) GetPromoByID(promo *models.Promo, promoID string) error {
	tx := p.db.First(promo, "id = ?", promoID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrPromoNotFound
	}

	return tx.Error
}

// ListAllPromos list all promos in the DB
func (p *PromoRepository) ListAllPromos(promos *[]models.Promo) error {
	return p.db.Find(promos).Error
}

// ListAllUsersByPromoID list all users of a specific Promo in the DB
func (p *PromoRepository) ListAllUsersByPromoID(promos *[]models.User, promoID string) error {
	return p.db.Find(promos, "promo_id=?", promoID).Error
}

// CreatePromo create a promo in the DB
func (p *PromoRepository) CreatePromo(promo *models.Promo) error {
	return p.db.Create(promo).Error
}

// UpdatePromo update a promo in the DB
func (p *PromoRepository) UpdatePromo(promo *models.Promo) error {
	return p.db.Save(promo).Error
}

// DeleteByPromoID delete a comment by ID in the DB
func (p *PromoRepository) DeleteByPromoID(promoID string) error {
	tx := p.db.Delete(&models.Promo{}, "id = ?", promoID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrPromoNotFound
	}

	return tx.Error
}

package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
)

// PromoHandler is a struct to define promo handler
type PromoHandler struct {
	repository *repository.PromoRepository
}

// NewPromoHandler is a factory promo handler
func NewPromoHandler(repository *repository.PromoRepository) *PromoHandler {
	return &PromoHandler{repository: repository}
}

// ListPromoUsers get users from the same promo
func (p *PromoHandler) ListPromoUsers(c *gin.Context) {
	id, _ := c.Params.Get("id")
	promo := &[]models.User{}

	err := p.repository.ListAllUsersByPromoID(promo, id)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, promo)

}

// ListPromos respond the list of all promos
func (p *PromoHandler) ListPromos(c *gin.Context) {
	promos := &[]models.Promo{}

	err := p.repository.ListAllPromos(promos)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, promos)

}

// CreatePromo create a promo
func (p *PromoHandler) CreatePromo(c *gin.Context) {
	promo := &models.Promo{}

	err := c.ShouldBindJSON(promo)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = p.repository.CreatePromo(promo)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, promo)

}

// DeletePromo delete a specific promo
func (p *PromoHandler) DeletePromo(c *gin.Context) {
	promoID, _ := c.Params.Get("id")

	err := p.repository.DeleteByPromoID(promoID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)

}

// UpdatePromo update a specific promo
func (p *PromoHandler) UpdatePromo(c *gin.Context) {
	promoID, _ := c.Params.Get("id")
	promo := &models.Promo{}

	err := p.repository.GetPromoByID(promo, promoID)
	if err != nil {
		if errors.Is(err, repository.ErrPromoNotFound) {
			httpError.NotFound(c, "Promo", promoID, err)
		} else {
			httpError.Internal(c, err)
		}
		return
	}

	err = c.ShouldBindJSON(promo)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = p.repository.UpdatePromo(promo)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, promo)
}

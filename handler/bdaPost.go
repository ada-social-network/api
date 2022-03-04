package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// BdaPostHandler is a struct to define bda post handler
type BdaPostHandler struct {
	repository *repository.BdaPostRepository
}

// NewBdaPostHandler is a factory for bda post handler
func NewBdaPostHandler(repository *repository.BdaPostRepository) *BdaPostHandler {
	return &BdaPostHandler{repository: repository}
}

// ListBdaPost respond a list of bda posts
func (bp *BdaPostHandler) ListBdaPost(c *gin.Context) {
	bdaPosts := &[]models.BdaPost{}

	err := bp.repository.ListAllBdaPosts(bdaPosts)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, bdaPosts)
}

// CreateBdaPost create a bda post
func (bp *BdaPostHandler) CreateBdaPost(c *gin.Context) {
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	bdaPost := &models.BdaPost{}

	err = c.ShouldBindJSON(bdaPost)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	bdaPost.UserID = user.ID

	err = bp.repository.CreateBdaPost(bdaPost)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, bdaPost)
}

// DeleteBdaPost delete a specific bda post
func (bp *BdaPostHandler) DeleteBdaPost(c *gin.Context) {
	id, _ := c.Params.Get("id")

	err := bp.repository.DeleteBdaPostByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrBdaPostNotFound) {
			httpError.NotFound(c, "comment", id, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

// GetBdaPost get a specific bda post
func (bp *BdaPostHandler) GetBdaPost(c *gin.Context) {
	id, _ := c.Params.Get("id")

	bdaPost := &models.BdaPost{}

	err := bp.repository.GetBdaPostByID(bdaPost, id)
	if err != nil {
		if errors.Is(err, repository.ErrBdaPostNotFound) {
			httpError.NotFound(c, "bdaPost", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, bdaPost)
}

// UpdateBdaPost update a specific bda post
func (bp *BdaPostHandler) UpdateBdaPost(c *gin.Context) {
	id, _ := c.Params.Get("id")
	bdaPost := &models.BdaPost{}

	err := bp.repository.GetBdaPostByID(bdaPost, id)
	if err != nil {
		if errors.Is(err, repository.ErrBdaPostNotFound) {
			httpError.NotFound(c, "bdaPost", id, err)

		}
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(bdaPost)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = bp.repository.UpdateBdaPost(bdaPost)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, bdaPost)
}

// LikeBdaPostResponse defines the Like response for a BdaPost
type LikeBdaPostResponse struct {
	models.Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	BdaPostID uuid.UUID `gorm:"type=uuid" json:"bdaPostId"`
}

// createBdaPostLikeResponse map the values of like to likeBdaPostResponse
func createBdaPostLikeResponse(like models.Like) LikeBdaPostResponse {
	return LikeBdaPostResponse{
		Base:      like.Base,
		UserID:    like.UserID,
		BdaPostID: like.BdaPostID,
	}
}

// CreateBdaPostLike create a like
func CreateBdaPostLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPostID, _ := c.Params.Get("id")

		like := &models.Like{}

		err = c.ShouldBindJSON(like)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPostUUID, err := uuid.FromString(bdaPostID)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		like.BdaPostID = bdaPostUUID
		like.UserID = user.ID

		tx := db.Where("user_id= ? AND bda_post_id= ?", like.UserID, like.BdaPostID).Find(like)
		if tx.Error != nil {
			httpError.Internal(c, tx.Error)
			return
		}

		if tx.RowsAffected > 0 {
			httpError.AlreadyLiked(c, "user_id", like.UserID.String())
			return
		}

		result := db.Create(like)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, createBdaPostLikeResponse(*like))
	}
}

// ListBdaPostLikes get likes of a bda post
func ListBdaPostLikes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		likes := &[]models.Like{}
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Find(likes, "bda_post_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		var liked = &models.Like{}
		tx := db.Where("user_id= ? AND bda_post_id= ?", user.ID, id).Find(liked)
		if tx.Error != nil {
			httpError.Internal(c, tx.Error)
			return
		}

		var isLikedByCurrentUser bool
		if tx.RowsAffected > 0 {
			isLikedByCurrentUser = true
		}

		likesResponse := []interface{}{}

		for _, like := range *likes {
			likesResponse = append(likesResponse, createBdaPostLikeResponse(like))
		}

		c.JSON(200, NewLikeCollection(likesResponse, isLikedByCurrentUser))
	}
}

// DeleteBdaPostLike delete a specific like
func DeleteBdaPostLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("likeId")

		result := db.Delete(&models.Like{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

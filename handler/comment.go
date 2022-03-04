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

// CommentHandler is a struct to define comment handler
type CommentHandler struct {
	repository *repository.CommentRepository
}

// NewCommentHandler is a factory comment handler
func NewCommentHandler(repository *repository.CommentRepository) *CommentHandler {
	return &CommentHandler{repository: repository}
}

// CreateBdaPostComment create a comment
func (co *CommentHandler) CreateBdaPostComment(c *gin.Context) {
	commentRequestURI := &models.CommentRequestURI{}
	commentRequest := &models.CommentRequest{}

	err := c.ShouldBindUri(commentRequestURI)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(commentRequest)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	comment, err := co.repository.CreateComment(user.ID, uuid.FromStringOrNil(commentRequestURI.BdaPostID), commentRequest.Content)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, models.CommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		BdaPostID: comment.BdaPostID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	})
}

// UpdateBdaPostComment update a specific comment
func (co *CommentHandler) UpdateBdaPostComment(c *gin.Context) {
	//can be c.Request.URL.Query().Get("id") but it's a shorter notation
	commentID, _ := c.Params.Get("commentId")
	comment := &models.Comment{}

	err := co.repository.GetCommentByID(comment, commentID)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			httpError.NotFound(c, "comment", commentID, err)

		}
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(comment)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = co.repository.UpdateComment(comment)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, comment)
}

// DeleteBdaPostComment delete a specific comment
func (co *CommentHandler) DeleteBdaPostComment(c *gin.Context) {
	//can be c.Request.URL.Query().Get("id") but it's a shorter notation
	commentID, _ := c.Params.Get("commentId")

	err := co.repository.DeleteByCommentID(commentID)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			httpError.NotFound(c, "comment", commentID, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

// GetBdaPostComment get a specific comment
func (co *CommentHandler) GetBdaPostComment(c *gin.Context) {
	//can be c.Request.URL.Query().Get("id") but it's a shorter notation
	commentID, _ := c.Params.Get("commentId")

	comment := &models.Comment{}

	err := co.repository.GetCommentByID(comment, commentID)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			httpError.NotFound(c, "comment", commentID, err)

		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, comment)
}

// ListBdaPostComments get comments of a bda post
func (co *CommentHandler) ListBdaPostComments(c *gin.Context) {
	id, _ := c.Params.Get("id")
	comments := &[]models.Comment{}

	err := co.repository.ListAllCommentsByBdaPostID(comments, id)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, comments)
}

// LikeCommentResponse defines the Like response for a BdaPost
type LikeCommentResponse struct {
	models.Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	CommentID uuid.UUID `gorm:"type=uuid" json:"commentId"`
}

// createCommentLikeResponse map the values of like to likeBdaPostResponse
func createCommentLikeResponse(like models.Like) LikeCommentResponse {
	return LikeCommentResponse{
		Base:      like.Base,
		UserID:    like.UserID,
		CommentID: like.CommentID,
	}
}

// CreateCommentLike create a like
func CreateCommentLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		commentID, _ := c.Params.Get("id")

		like := &models.Like{}

		err = c.ShouldBindJSON(like)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		commentUUID, err := uuid.FromString(commentID)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		like.CommentID = commentUUID
		like.UserID = user.ID

		tx := db.Where("user_id= ? AND comment_id= ?", like.UserID, like.CommentID).Find(like)
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

		c.JSON(200, createCommentLikeResponse(*like))
	}
}

// ListCommentLikes get likes of a bda post
func ListCommentLikes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		likes := &[]models.Like{}
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Find(likes, "comment_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		var liked = &models.Like{}
		tx := db.Where("user_id= ? AND comment_id= ?", user.ID, id).Find(liked)
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
			likesResponse = append(likesResponse, createCommentLikeResponse(like))
		}

		c.JSON(200, NewLikeCollection(likesResponse, isLikedByCurrentUser))
	}
}

// DeleteCommentLike delete a specific like
func DeleteCommentLike(db *gorm.DB) gin.HandlerFunc {
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

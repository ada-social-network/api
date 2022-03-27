package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	bdapostID, _ := c.Params.Get("id")

	bdaPostUUID, err := uuid.FromString(bdapostID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	comment := &models.Comment{}
	err = c.ShouldBindJSON(comment)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	comment.BdaPostID = bdaPostUUID
	comment.UserID = user.ID

	err = co.repository.CreateComment(comment)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, comment)
}

// UpdateBdaPostComment update a specific comment
func (co *CommentHandler) UpdateBdaPostComment(c *gin.Context) {
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
	commentID, _ := c.Params.Get("commentId")

	err := co.repository.DeleteCommentByID(commentID)
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
func (co *CommentHandler) CreateCommentLike(c *gin.Context) {
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

	exist, err := co.repository.CheckLikeByUserAndCommentID(like, like.UserID, like.CommentID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	if exist {
		httpError.AlreadyLiked(c, "user_id", like.UserID.String())
		return
	}

	err = co.repository.CreateLike(like)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, createCommentLikeResponse(*like))
}

// ListCommentLikes get likes of a bda post
func (co *CommentHandler) ListCommentLikes(c *gin.Context) {
	commentID, _ := c.Params.Get("id")
	likes := &[]models.Like{}

	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = co.repository.ListAllPostsByCommentID(likes, commentID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	var liked = &models.Like{}

	exist, err := co.repository.CheckLikeByUserAndCommentID(liked, user.ID, uuid.FromStringOrNil(commentID))
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	var isLikedByCurrentUser bool
	if exist {
		isLikedByCurrentUser = true
	}

	likesResponse := []interface{}{}

	for _, like := range *likes {
		likesResponse = append(likesResponse, createCommentLikeResponse(like))
	}

	c.JSON(200, NewLikeCollection(likesResponse, isLikedByCurrentUser))
}

// DeleteCommentLike delete a specific like
func (co *CommentHandler) DeleteCommentLike(c *gin.Context) {
	id, _ := c.Params.Get("likeId")

	err := co.repository.DeleteLikeByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrLikeNotFound) {
			httpError.NotFound(c, "like", id, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

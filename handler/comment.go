package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// CreateBdaPostComment create a comment
func CreateBdaPostComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdapostID, _ := c.Params.Get("id")

		comment := &models.Comment{}
		err = c.ShouldBindJSON(comment)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPostUUID, err := uuid.FromString(bdapostID)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		comment.BdaPostID = bdaPostUUID
		comment.UserID = user.ID

		result := db.Create(comment)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, comment)
	}
}

// UpdateBdaPostComment update a specific comment
func UpdateBdaPostComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		commentID, _ := c.Params.Get("commentId")
		comment := &models.Comment{}

		result := db.First(comment, "id = ?", commentID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Comment", commentID, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(comment)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(comment)

		c.JSON(200, comment)
	}
}

// DeleteBdaPostComment delete a specific comment
func DeleteBdaPostComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		commentID, _ := c.Params.Get("commentId")

		result := db.Delete(&models.Comment{}, "id = ?", commentID)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// GetBdaPostComment get a specific comment
func GetBdaPostComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
		commentID, _ := c.Params.Get("commentId")

		comment := &models.Comment{}

		result := db.First(comment, "id = ?", commentID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "Comment", commentID, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, comment)
	}
}

// ListBdaPostComments get comments of a bda post
func ListBdaPostComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		comments := &[]models.Comment{}

		result := db.Find(comments, "bda_post_id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, comments)
	}

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

		result := db.Find(likes, "comment_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		likesResponse := []interface{}{}

		for _, like := range *likes {
			likesResponse = append(likesResponse, createCommentLikeResponse(like))
		}

		c.JSON(200, NewCollection(likesResponse))
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

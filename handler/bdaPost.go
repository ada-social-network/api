package handler

import (
	"errors"
	"time"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ListBdaPost respond a list of bda posts
func ListBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bdaPosts := &[]models.BdaPost{}

		result := db.Find(&bdaPosts)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, bdaPosts)
	}
}

// CreateBdaPost create a bda post
func CreateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		result := db.Create(bdaPost)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, bdaPost)
	}
}

// DeleteBdaPost delete a specific bda post
func DeleteBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		result := db.Delete(&models.BdaPost{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

// BdaPostResponse defines the bdaPost response for the method get bdaPost
type BdaPostResponse struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`

	Comments []models.Comment `json:"comments"`
	Likes    []models.Like    `json:"likes"`

	IsLikedByCurrentUser bool      `json:"isLikedByCurrentUser"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

// CreateBdaPostResponse map the values of bdaPost, likes, comments and user on a BdaPostResponse and check if the bdaPoost is likedByCurrentUser
func CreateBdaPostResponse(bdaPost *models.BdaPost, likes []models.Like, comments []models.Comment, currentUser *models.User) BdaPostResponse {
	isLikedByCurrentUser := false

	for _, like := range bdaPost.Likes {
		if like.UserID == currentUser.ID {
			isLikedByCurrentUser = true
		}
	}

	return BdaPostResponse{
		ID:                   bdaPost.ID,
		Title:                bdaPost.Title,
		Content:              bdaPost.Content,
		UserID:               bdaPost.UserID,
		Comments:             comments,
		Likes:                likes,
		IsLikedByCurrentUser: isLikedByCurrentUser,
		CreatedAt:            bdaPost.CreatedAt,
		UpdatedAt:            bdaPost.UpdatedAt,
	}
}

// GetBdaPost get a specific bda post
func GetBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		result := db.Preload("Likes").Preload("Comments").First(bdaPost, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "BdaPost", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		c.JSON(200, CreateBdaPostResponse(bdaPost, bdaPost.Likes, bdaPost.Comments, user))
	}
}

// UpdateBdaPost update a specific bda post
func UpdateBdaPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		bdaPost := &models.BdaPost{}

		result := db.First(bdaPost, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpError.NotFound(c, "BdaPost", id, result.Error)
			} else {
				httpError.Internal(c, result.Error)
			}
			return
		}

		err := c.ShouldBindJSON(bdaPost)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		db.Save(bdaPost)

		c.JSON(200, bdaPost)
	}
}

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

// LikeBdaPostResponse defines the Like response for a BdaPost
type LikeBdaPostResponse struct {
	models.Base
	UserID    uuid.UUID `gorm:"type=uuid" json:"userId" `
	BdaPostID uuid.UUID `gorm:"type=uuid" json:"bdapostId"`
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

		result := db.Find(likes, "bda_post_id= ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		likesResponse := []interface{}{}

		for _, like := range *likes {
			likesResponse = append(likesResponse, createBdaPostLikeResponse(like))
		}

		c.JSON(200, NewCollection(likesResponse))
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

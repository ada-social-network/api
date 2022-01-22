package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
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

// GetBdaPost get a specific bda post
func GetBdaPost(db *gorm.DB) gin.HandlerFunc {
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

		c.JSON(200, bdaPost)
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

// CreateBdaPostLike create a like
func CreateBdaPostLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetCurrentUser(c)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdapostID, _ := c.Params.Get("id")

		like := &models.Like{}

		err = c.ShouldBindJSON(like)
		if err != nil {
			httpError.Internal(c, err)
			return
		}

		bdaPostUUID, err := uuid.FromString(bdapostID)
		if err != nil {
			httpError.Internal(c, err)
			return
		}
		like.BdaPostID = bdaPostUUID
		like.UserID = user.ID

		result := db.Create(like)
		if result.Error != nil {
			httpError.Internal(c, err)
			return
		}

		c.JSON(200, like)
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

		count := result.RowsAffected

		c.JSON(200, models.CountAndLikes{Likes: likes, Count: int(count)})
	}
}

// DeleteBdaPostLike delete a specific like
func DeleteBdaPostLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("likeId")

		result := db.Delete(&models.Post{}, "id = ?", id)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(204, nil)
	}
}

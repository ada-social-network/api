package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// PostHandler is a struct to define post handler
type PostHandler struct {
	repository *repository.PostRepository
}

// NewPostHandler is a factory post handler
func NewPostHandler(repository *repository.PostRepository) *PostHandler {
	return &PostHandler{repository: repository}
}

// ListPost get posts of a topic
func (p *PostHandler) ListPost(c *gin.Context) {
	id, _ := c.Params.Get("id")
	posts := &[]models.Post{}

	err := p.repository.ListAllPostsByTopicID(posts, id)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, posts)
}

// CreatePost create a post
func (p *PostHandler) CreatePost(c *gin.Context) {
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	topicID, _ := c.Params.Get("id")

	topicUUID, err := uuid.FromString(topicID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	post := &models.Post{}
	err = c.ShouldBindJSON(post)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	post.UserID = user.ID
	post.TopicID = topicUUID

	err = p.repository.CreatePost(post)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, post)
}

// DeletePost delete a specific post
func (p *PostHandler) DeletePost(c *gin.Context) {
	postID, _ := c.Params.Get("postId")

	err := p.repository.DeletePostByID(postID)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			httpError.NotFound(c, "post", postID, err)
			return
		}

		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)
}

// GetPost get a specific post
func (p *PostHandler) GetPost(c *gin.Context) {
	postID, _ := c.Params.Get("postID")

	post := &models.Post{}

	err := p.repository.GetPostByID(post, postID)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			httpError.NotFound(c, "post", postID, err)
		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, post)
}

// UpdatePost update a specific post
func (p *PostHandler) UpdatePost(c *gin.Context) {
	postID, _ := c.Params.Get("postID")
	post := &models.Post{}

	err := p.repository.GetPostByID(post, postID)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			httpError.NotFound(c, "post", postID, err)
		}
		httpError.Internal(c, err)
		return
	}

	err = c.ShouldBindJSON(post)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = p.repository.UpdatePost(post)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, post)
}

// LikeHandler is a struct to define like handler
type LikeHandler struct {
	repository *repository.LikeRepository
}

// NewLikeHandler is a factory like handler
func NewLikeHandler(repository *repository.LikeRepository) *LikeHandler {
	return &LikeHandler{repository: repository}
}

// LikePostResponse defines the Like response for a Post
type LikePostResponse struct {
	models.Base
	UserID uuid.UUID `gorm:"type=uuid" json:"userId" `
	PostID uuid.UUID `gorm:"type=uuid" json:"postId"`
}

// createPostLikeResponse map the values of like to likePostResponse
func createPostLikeResponse(like models.Like) LikePostResponse {
	return LikePostResponse{
		Base:   like.Base,
		UserID: like.UserID,
		PostID: like.PostID,
	}
}

// CreatePostLike create a like
func (l *LikeHandler) CreatePostLike(c *gin.Context) {
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	postID, _ := c.Params.Get("id")

	like := &models.Like{}

	err = c.ShouldBindJSON(like)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	PostUUID, err := uuid.FromString(postID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}
	like.PostID = PostUUID
	like.UserID = user.ID

	exist, err := l.repository.CheckLikeByUserAndPostID(like, like.UserID, like.PostID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}
	if exist {
		httpError.AlreadyLiked(c, "user_id", like.UserID.String())
		return
	}

	err = l.repository.CreateLike(like)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, createPostLikeResponse(*like))
}

// ListPostLikes get likes of a post
func (l *LikeHandler) ListPostLikes(c *gin.Context) {
	postID, _ := c.Params.Get("id")
	likes := &[]models.Like{}

	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = l.repository.ListAllPostsByPostID(likes, postID)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	var liked = &models.Like{}

	exist, err := l.repository.CheckLikeByUserAndPostID(liked, user.ID, uuid.FromStringOrNil(postID))
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
		likesResponse = append(likesResponse, createPostLikeResponse(like))
	}

	c.JSON(200, NewLikeCollection(likesResponse, isLikedByCurrentUser))
}

// DeletePostLike delete a specific like
func (l *LikeHandler) DeletePostLike(c *gin.Context) {
	id, _ := c.Params.Get("likeId")

	err := l.repository.DeleteLikeByID(id)
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

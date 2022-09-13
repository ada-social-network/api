package handler

import (
	"errors"

	httpError "github.com/ada-social-network/api/error"
	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserHandler is a struct to define user handler
type UserHandler struct {
	repository *repository.UserRepository
}

// NewUserHandler is a factory user handler
func NewUserHandler(repository *repository.UserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}

// UserResponse define a user response
type UserResponse struct {
	models.Base
	LastName       string           `json:"lastName" binding:"required,min=2,max=20"`
	FirstName      string           `json:"firstName" binding:"required,min=2,max=20"`
	Email          string           `json:"email" binding:"required,email" gorm:"unique"`
	Password       string           `json:"password"`
	DateOfBirth    string           `json:"dateOfBirth"`
	Apprenticeship string           `json:"apprenticeAt"`
	ProfilPic      string           `json:"profilPic"`
	Biography      string           `json:"biography"`
	CoverPic       string           `json:"coverPic"`
	PrivateMail    string           `json:"privateMail"`
	ProjectPerso   string           `json:"projectPerso"`
	ProjectPro     string           `json:"projectPro"`
	Instagram      string           `json:"instagram"`
	Facebook       string           `json:"facebook"`
	Github         string           `json:"github"`
	Linkedin       string           `json:"linkedin"`
	MBTI           string           `json:"mbti"`
	Admin          bool             `json:"isAdmin"`
	PromoID        uuid.UUID        `gorm:"type=uuid" json:"promoId"`
	BdaPosts       []models.BdaPost `json:"bdaPosts"`
	Posts          []models.Post    `json:"posts"`
	Comments       []models.Comment `json:"comments"`
	Topics         []models.Topic   `json:"topics"`
	Likes          []models.Like    `json:"likes"`
}

// UpdatePasswordRequest is the request for the password change
type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

type userRegister struct {
	LastName  string `json:"lastName" binding:"required,min=2,max=20"`
	FirstName string `json:"firstName" binding:"required,min=2,max=20"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=32"`
}

// Register register a user
func (us *UserHandler) Register(c *gin.Context) {
	userRegister := &userRegister{}

	err := c.BindJSON(userRegister)
	if err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			httpError.Validation(c, ve)
			return
		}

		httpError.Internal(c, err)
		return
	}

	user := &models.User{
		LastName:  userRegister.LastName,
		FirstName: userRegister.FirstName,
		Email:     userRegister.Email,
		Password:  userRegister.Password,
	}

	exist, err := us.repository.CheckUniqueMailInUsers(user, user.Email)
	if err != nil {
		httpError.Internal(c, err)
		return
	}
	if exist {
		httpError.AlreadyExist(c, "email", user.Email)
		return
	}

	err = us.repository.CreateUserWithPassword(user, user.Password)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, createUserResponse(user))
}

// Me provide informations about the connected user
func (us *UserHandler) Me(c *gin.Context) {
	user, exist := c.Get("id")
	if !exist {
		httpError.Internal(c, errors.New("jwt does not exist"))
		return
	}

	u, ok := user.(*models.User)
	if !ok {
		httpError.Internal(c, errors.New("this is not a user"))
		return
	}

	err := us.repository.GetUserByID(u, u.ID.String())
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, createUserResponse(u))
}

// UpdatePassword update password of the current user
func (us *UserHandler) UpdatePassword(c *gin.Context) {
	user, err := GetCurrentUser(c)
	if err != nil {
		httpError.Internal(c, err)
	}

	updatePasswordRequest := &UpdatePasswordRequest{}
	err = c.ShouldBindJSON(updatePasswordRequest)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = us.repository.UpdateUserWithPassword(user, updatePasswordRequest.Password)
	if err != nil {
		httpError.Internal(c, err)
	}

	c.JSON(204, nil)
}

// ListUser respond a list of users
func (us *UserHandler) ListUser(c *gin.Context) {
	users := &[]models.User{}

	err := us.repository.ListAllUser(users)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, createUsersResponse(users))

}

// CreateUser create a user
func (us *UserHandler) CreateUser(c *gin.Context) {
	user := &models.User{}

	err := c.ShouldBindJSON(user)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	err = us.repository.CreateUserWithPassword(user, user.Password)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, user)
}

// DeleteUser delete a specific user
func (us *UserHandler) DeleteUser(c *gin.Context) {

	user, _ := c.Params.Get("id")

	err := us.repository.DeleteByUserID(user)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	c.JSON(204, nil)

}

// GetUser get a specific user
func (us *UserHandler) GetUser(c *gin.Context) {
	userID, _ := c.Params.Get("id")

	user := &models.User{}

	err := us.repository.GetUserByID(user, userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			httpError.NotFound(c, "user", userID, err)

		}
		httpError.Internal(c, err)
		return
	}

	c.JSON(200, user)
}

// UpdateUser update a specific user
func (us *UserHandler) UpdateUser(c *gin.Context) {
	userID, _ := c.Params.Get("id")
	user := &models.User{}

	err := us.repository.GetUserByID(user, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpError.NotFound(c, "User", userID, err)
		} else {
			httpError.Internal(c, err)
		}
		return
	}

	err = c.ShouldBindJSON(user)
	if err != nil {
		httpError.Internal(c, err)
		return
	}

	//if user.Password == "" {
	err = us.repository.UpdateUserWithoutPassword(user)
	//fmt.Println("is", user.FirstName, user.Password)
	//} else {
	//err = us.repository.UpdateUserWithPassword(user, user.Password)
	//mt.Println("coucou", user.FirstName, user.Password)
	//}

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			httpError.NotFound(c, "User", userID, err)
		} else {
			httpError.Internal(c, err)
		}
		return
	}

	c.JSON(200, createUserResponse(user))
}

// createUserResponse map the values of user to createUserResponse
func createUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Base:           user.Base,
		LastName:       user.LastName,
		FirstName:      user.FirstName,
		Email:          user.Email,
		DateOfBirth:    user.DateOfBirth,
		Apprenticeship: user.Apprenticeship,
		ProfilPic:      user.ProfilPic,
		Biography:      user.Biography,
		CoverPic:       user.CoverPic,
		PrivateMail:    user.PrivateMail,
		ProjectPerso:   user.ProjectPerso,
		ProjectPro:     user.ProjectPro,
		Instagram:      user.Instagram,
		Facebook:       user.Facebook,
		Github:         user.Github,
		Linkedin:       user.Linkedin,
		MBTI:           user.MBTI,
		Admin:          user.Admin,
		PromoID:        user.PromoID,
		BdaPosts:       user.BdaPosts,
		Posts:          user.Posts,
		Comments:       user.Comments,
		Topics:         user.Topics,
		Likes:          user.Likes,
	}
}

// createUserResponse map the values of all users to a list of createUserResponse
func createUsersResponse(users *[]models.User) []UserResponse {
	usersList := []UserResponse{}
	for _, u := range *users {
		usersList = append(usersList, createUserResponse(&u))
	}

	return usersList
}

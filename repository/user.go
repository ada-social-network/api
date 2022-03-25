package repository

import (
	"errors"

	"github.com/ada-social-network/api/models"
	"gorm.io/gorm"
)

// ErrUserNotFound is an error when resource is not found
var (
	ErrUserNotFound     = errors.New("user not found")
	ErrPasswordNotFound = errors.New("password not found")
)

// UserRepository is a repository for user resource
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is to create a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser create a user in the DB
func (us *UserRepository) CreateUser(user *models.User) error {
	return us.db.Create(user).Error
}

// GetUserByID get a user by id in the DB
func (us *UserRepository) GetUserByID(user *models.User, userID string) error {
	tx := us.db.First(user, "id = ?", userID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}

	return tx.Error
}

// ListAllUser list all users in the DB
func (us *UserRepository) ListAllUser(users *[]models.User) error {
	return us.db.Find(users).Error
}

// UpdateUserWithoutPassword update a user in the DB without password
func (us *UserRepository) UpdateUserWithoutPassword(user *models.User) error {
	return us.db.Omit("Password").Save(user).Error
}

// UpdateUserWithPassword update user password only
func (us *UserRepository) UpdateUserWithPassword(user *models.User, password string) error {
	passwordEncrypted, err := models.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = passwordEncrypted
	return us.db.Save(user).Error
}

// DeleteByUserID delete a comment by ID in the DB
func (us *UserRepository) DeleteByUserID(userID string) error {
	tx := us.db.Delete(&models.User{}, "id = ?", userID)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}

	return tx.Error
}

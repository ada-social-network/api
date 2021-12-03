package models

// Post define post resource
type Post struct {
	Base
	Content string `json:"content" binding:"required,min=4,max=21474"`

	// By default, gorm will try to use UserID as a foreign key to the model User
	UserID uint `json:"user_id" binding:"required"`
}

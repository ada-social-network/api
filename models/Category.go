package models

// Category model
type Category struct {
	Base
	Name   string  `json:"name" binding:"required"`
	Topics []Topic `json:"topics"`
}

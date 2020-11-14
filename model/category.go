package model

import "github.com/jinzhu/gorm"

// Category of a book
type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	BookID      int    `sql:"index"`
}

// CategoryRequest request body
type CategoryRequest struct {
	Name        string
	Description string
	BookIDs     []int
}

// CategoryResponse response
type CategoryResponse struct {
	CategoryID int
}

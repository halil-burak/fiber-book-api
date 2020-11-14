package model

import "github.com/jinzhu/gorm"

// Author of a book
type Author struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

// AuthorRequest request body
type AuthorRequest struct {
	Name     string
	LastName string
}

// AuthorResponse response
type AuthorResponse struct {
	AuthorID int
}

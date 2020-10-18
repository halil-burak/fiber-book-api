package model

import "github.com/jinzhu/gorm"

// Category of a book
type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

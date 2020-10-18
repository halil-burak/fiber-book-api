package model

import "github.com/jinzhu/gorm"

// Author of a book
type Author struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

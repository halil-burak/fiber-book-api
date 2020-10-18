package model

import "github.com/jinzhu/gorm"

// Book has a title, an author and a rating
type Book struct {
	gorm.Model
	Title      string     `json:"title"`
	Author     Author     `json:"author"`
	Rating     int        `json:"rating"`
	Categories []Category `gorm:"many2many:book_categories;"`
}

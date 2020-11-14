package model

import "github.com/jinzhu/gorm"

const (
	// AssociationCategories is for many to many relation
	AssociationCategories = "Categories"
)

// Book has a title, an author and a rating
type Book struct {
	gorm.Model
	Title      string `json:"title"`
	AuthorID   int
	Author     Author     `json:"author"`
	Rating     int        `json:"rating"`
	Categories []Category `gorm:"many2many:book_categories json: 'categories'"`
}

// BookRequest request body
type BookRequest struct {
	Title         string
	AuthorID      int
	Rating        int
	CategoryNames []string
}

// BookResponse response
type BookResponse struct {
	BookID uint
}

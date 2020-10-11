package book

import (
	"github.com/gofiber/fiber"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/jinzhu/gorm"
)

// Book has a title, an author and a rating
type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks returns all books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook returns single book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// NewBook adds a new book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(book)
	c.JSON(book)
}

// DeleteBook removes a book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted.")
}

// UpdateBook updates an existing book
func UpdateBook(c *fiber.Ctx) {
	c.Send("Updates a book")
}

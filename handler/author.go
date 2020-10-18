package handler

import (
	"github.com/gofiber/fiber"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/halil-burak/fiber-rest-api/model"
)

// NewAuthor adds a new author
func NewAuthor(c *fiber.Ctx) {
	db := database.DBConn
	author := new(model.Author)
	db.Create(author)
}

// GetAuthors returns all authors
func GetAuthors(c *fiber.Ctx) {
	db := database.DBConn
	var authors []model.Author
	db.Find(&authors)
	c.JSON(authors)
}

// GetAuthor returns single author
func GetAuthor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var author model.Author
	db.First(&author, id)
	c.JSON(author)
}

// DeleteAuthor deletes an author
func DeleteAuthor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var author model.Author
	db.First(&author, id)

	if author.ID == 0 {
		c.Status(404).Send("No book found with ID")
		return
	}

	db.Delete(&author)
	c.Send("Author removed")
}

// UpdateAuthor updates an author
func UpdateAuthor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var oldAuthor model.Author
	db.First(&oldAuthor, id)

	if oldAuthor.ID == 0 {
		c.Status(404).Send("No author found with ID")
		return
	}

	author := new(model.Author)
	if err := c.BodyParser(author); err != nil {
		c.Status(503).Send(err)
		return
	}

	oldAuthor.Name = author.Name
	oldAuthor.LastName = author.LastName
	db.Save(&oldAuthor)
	c.Send("Updated author.")
}

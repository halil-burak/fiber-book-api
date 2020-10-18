package handler

import (
	"github.com/gofiber/fiber"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/halil-burak/fiber-rest-api/model"
)

// NewCategory adds a new category
func NewCategory(c *fiber.Ctx) {
	db := database.DBConn
	category := new(model.Category)
	db.Create(&category)
}

// GetCategories returns all categories
func GetCategories(c *fiber.Ctx) {
	db := database.DBConn
	var categories []model.Category
	db.Find(&categories)
	c.JSON(categories)
}

// GetCategory returns single category
func GetCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var category model.Category
	db.First(&category, id)
	c.JSON(category)
}

// DeleteCategory deletes a category
func DeleteCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var category model.Category
	db.First(&category, id)

	if category.ID == 0 {
		c.Status(404).Send("No category found with ID")
		return
	}

	db.Delete(&category)
	c.Send("Category removed")
}

// UpdateCategory updates a category
func UpdateCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var oldCategory model.Category
	db.First(&oldCategory, id)

	if oldCategory.ID == 0 {
		c.Status(404).Send("No category found with ID")
		return
	}

	category := new(model.Category)
	if err := c.BodyParser(category); err != nil {
		c.Status(503).Send(err)
		return
	}

	oldCategory.Name = category.Name
	oldCategory.Description = category.Description
	db.Save(&oldCategory)
	c.Send("Updated category.")
}

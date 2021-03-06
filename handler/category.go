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

	if err := c.BodyParser(category); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&category)
	c.JSON(category)
}

// CreateIfNotExists creates a category if the the categoryName does not exist
func CreateIfNotExists(categoryName string) (*model.Category, error) {
	db := database.DBConn
	var category model.Category
	response := db.FirstOrCreate(&category, model.Category{Name: categoryName})
	if response.Error != nil {
		return nil, response.Error
	}
	return &category, nil
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

	if category.ID == 0 || category.Name == "" {
		c.Status(404).Send("No category found with ID")
		return
	}

	c.JSON(category)
}

// DeleteCategory deletes a category
func DeleteCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var category model.Category
	db.First(&category, id)

	if category.ID == 0 || category.Name == "" {
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

	if oldCategory.ID == 0 || oldCategory.Name == "" {
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

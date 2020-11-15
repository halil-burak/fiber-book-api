package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/halil-burak/fiber-rest-api/database"
	"github.com/halil-burak/fiber-rest-api/handler"
	"github.com/halil-burak/fiber-rest-api/hello"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send(hello.Hello())
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())
	app.Get("/", helloWorld)

	books := api.Group("/books")
	books.Get("/", handler.GetBooks)
	books.Get("/:id", handler.GetBook)
	books.Post("/", handler.NewBook)
	books.Delete("/:id", handler.DeleteBook)
	books.Put("/:id", handler.UpdateBook)
	books.Get("/:id/author", handler.GetAuthorOfBook)
	books.Get("/:id/categories", handler.GetCategoriesOfBook)

	authors := api.Group("/authors")
	authors.Post("/", handler.NewAuthor)
	authors.Get("/", handler.GetAuthors)
	authors.Get("/:id", handler.GetAuthor)
	authors.Delete("/:id", handler.DeleteAuthor)
	authors.Put("/:id", handler.UpdateAuthor)

	categories := api.Group("/categories")
	categories.Get("/", handler.GetCategories)
	categories.Get("/:id", handler.GetCategory)
	categories.Post("/", handler.NewCategory)
	categories.Delete("/:id", handler.DeleteCategory)
	categories.Put("/:id", handler.UpdateCategory)

	userLangs := api.Group("/users")
	userLangs.Get("/", handler.GetAllUsers)
	userLangs.Get("/:id", handler.GetOneUser)
	userLangs.Post("/", handler.AddUser)
	userLangs.Put("/:id", handler.UpdateUser)
}

func main() {
	app := fiber.New()
	database.ConnectDB()

	setupRoutes(app)
	app.Listen(3000)
}

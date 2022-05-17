package routes

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New() *fiber.App {
	app := fiber.New()

	// app.Use(cors.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/books", handlers.GetAllBooks)
	v1.Get("/books/:id", handlers.GetBookByID)
	v1.Post("/books", handlers.RegisterBook)
	v1.Delete("/books/:id", handlers.DeleteBook)

	return app
}

package routes

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func bookRoutes(v1 fiber.Router) {
	v1.Get("/books", handlers.GetAll)
	v1.Get("/books/:id", handlers.GetByID)
	v1.Post("/books", handlers.Create)
	v1.Delete("/books/:id", handlers.Delete)
}

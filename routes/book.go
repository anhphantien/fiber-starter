package routes

import (
	"fiber-starter/handlers"
	"fiber-starter/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(v1 fiber.Router) {
	v1.Use(middlewares.JwtAuth())

	v1.Get("books", handlers.BookHandler{}.GetAll)
	v1.Get("books/:id", handlers.BookHandler{}.GetByID)
	v1.Post("books", handlers.BookHandler{}.Create)
	v1.Delete("books/:id", handlers.BookHandler{}.Delete)
}

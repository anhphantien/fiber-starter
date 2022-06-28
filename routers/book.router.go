package routers

import (
	"fiber-starter/handlers"
	"fiber-starter/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(r fiber.Router) {
	// r.Use(
	// 	middlewares.JwtAuth(),
	// 	middlewares.AdminRole,
	// 	middlewares.UserRole,
	// )

	r.Get("books", handlers.BookHandler{}.GetList)

	r.Get("books/:id", handlers.BookHandler{}.GetByID)

	r.Post("books", handlers.BookHandler{}.Create)

	r.Put("books/:id", handlers.BookHandler{}.Update)

	r.Delete("books/:id",
		middlewares.JwtAuth(),
		// middlewares.AdminRole,
		middlewares.UserRole,
		handlers.BookHandler{}.Delete,
	)
}

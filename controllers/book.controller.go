package controllers

import (
	"fiber-starter/middlewares"
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func BookController(r fiber.Router) {
	// r.Use(
	// 	middlewares.JwtAuth(),
	// 	middlewares.AdminRole,
	// 	middlewares.UserRole,
	// )

	r.Get("books", services.BookService{}.GetList)

	r.Get("books/:id", services.BookService{}.GetByID)

	r.Post("books", services.BookService{}.Create)

	r.Put("books/:id", services.BookService{}.Update)

	r.Delete("books/:id",
		middlewares.JwtAuth(),
		// middlewares.AdminRole,
		middlewares.UserRole,
		services.BookService{}.Delete,
	)
}

package controllers

import (
	"fiber-starter/middlewares"
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func BookController(r fiber.Router) {
	// v1.Use(middlewares.JwtAuth())

	r.Get("books", services.BookService{}.GetAll)
	r.Get("books/:id", services.BookService{}.GetByID)
	r.Post("books", services.BookService{}.Create)
	r.Put("books/:id", services.BookService{}.Update)
	r.Delete("books/:id",
		middlewares.JwtAuth(),
		middlewares.AdminRole,
		middlewares.UserRole,
		services.BookService{}.Delete,
	)
}

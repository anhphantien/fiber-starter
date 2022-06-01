package controllers

import (
	"fiber-starter/middlewares"
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func BookController(v1 fiber.Router) {
	// v1.Use(middlewares.JwtAuth())

	v1.Get("books", services.BookService{}.GetAll)
	v1.Get("books/:id", services.BookService{}.GetByID)
	v1.Post("books", services.BookService{}.Create)
	v1.Put("books/:id", services.BookService{}.Update)
	v1.Delete("books/:id", middlewares.JwtAuth(), services.BookService{}.Delete)
}

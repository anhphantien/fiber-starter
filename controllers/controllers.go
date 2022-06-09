package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New() *fiber.App {
	app := fiber.New()

	// app.Use(cors.New())

	app.Get("swagger/*", swagger.HandlerDefault)

	v1 := app.Group("api/v1")
	AuthController(v1)
	BookController(v1)

	return app
}

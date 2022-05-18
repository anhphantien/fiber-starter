package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New() *fiber.App {
	app := fiber.New()

	// app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	bookRoutes(v1)

	return app
}

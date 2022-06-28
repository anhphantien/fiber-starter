package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	})

	// app.Use(cors.New())

	app.Get("swagger/*", swagger.HandlerDefault)

	v1 := app.Group("api/v1")
	AuthRouter(v1)
	BookRouter(v1)
	FileRouter(v1)
	UserRouter(v1)

	return app
}

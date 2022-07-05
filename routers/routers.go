package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	})
	app.Get("swagger/*", swagger.HandlerDefault)
	apiGroup(app, "api/v1")
	return app
}

func apiGroup(app *fiber.App, prefix string) {
	r := app.Group(prefix)
	AuthRouter(r)
	BookRouter(r)
	FileRouter(r)
	UserRouter(r)
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()

	app.Get("/", hello)

	app.Get("/swagger/", fiberSwagger.WrapHandler)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

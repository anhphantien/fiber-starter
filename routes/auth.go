package routes

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(v1 fiber.Router) {
	v1.Group("auth")

	v1.Post("auth/signin", handlers.AuthHandler{}.SignIn)
}

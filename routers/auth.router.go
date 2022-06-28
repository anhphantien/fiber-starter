package routers

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r fiber.Router) {
	r.Post("auth/login", handlers.AuthHandler{}.Login)
}

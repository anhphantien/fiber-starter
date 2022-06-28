package routers

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(r fiber.Router) {
	r.Get("users", handlers.UserHandler{}.GetList)
}

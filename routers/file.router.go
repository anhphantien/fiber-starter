package routers

import (
	"fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func FileRouter(r fiber.Router) {
	r.Post("file/upload", handlers.FileHandler{}.Upload)
}

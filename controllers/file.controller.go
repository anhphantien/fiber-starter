package controllers

import (
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func FileController(r fiber.Router) {
	r.Post("file/upload", services.FileService{}.Upload)
}

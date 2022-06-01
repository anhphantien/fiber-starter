package controllers

import (
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func AuthController(v1 fiber.Router) {
	v1.Post("auth/login", services.AuthService{}.Login)
}

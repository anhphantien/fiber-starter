package controllers

import (
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func AuthController(r fiber.Router) {
	r.Post("auth/login", services.AuthService{}.Login)
}

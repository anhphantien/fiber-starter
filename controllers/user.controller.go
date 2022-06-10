package controllers

import (
	"fiber-starter/services"

	"github.com/gofiber/fiber/v2"
)

func UserController(r fiber.Router) {
	r.Get("users", services.UserService{}.GetList)
}

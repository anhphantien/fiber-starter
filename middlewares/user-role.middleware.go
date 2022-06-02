package middlewares

import (
	"fiber-starter/config"

	"github.com/gofiber/fiber/v2"
)

func UserRole(c *fiber.Ctx) error {
	c.Locals(config.USER_ROLE, "USER")
	return c.Next()
}

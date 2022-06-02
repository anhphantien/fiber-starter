package middlewares

import (
	"fiber-starter/config"

	"github.com/gofiber/fiber/v2"
)

func AdminRole(c *fiber.Ctx) error {
	c.Locals(config.ADMIN_ROLE, "ADMIN")
	return c.Next()
}

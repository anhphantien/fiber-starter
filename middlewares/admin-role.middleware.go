package middlewares

import (
	"fiber-starter/config"
	"fiber-starter/enums"

	"github.com/gofiber/fiber/v2"
)

func AdminRole(c *fiber.Ctx) error {
	c.Locals(config.ADMIN_ROLE, enums.UserRole.ADMIN)
	return c.Next()
}

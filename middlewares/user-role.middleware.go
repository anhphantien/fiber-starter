package middlewares

import (
	"fiber-starter/config"
	"fiber-starter/enums"

	"github.com/gofiber/fiber/v2"
)

func UserRole(c *fiber.Ctx) error {
	c.Locals(config.USER_ROLE, enums.UserRole.USER)
	return c.Next()
}

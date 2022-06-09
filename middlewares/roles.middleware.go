package middlewares

import (
	"fiber-starter/enums"

	"github.com/gofiber/fiber/v2"
)

const (
	ADMIN_ROLE = "ADMIN_ROLE"
	USER_ROLE  = "USER_ROLE"
)

func AdminRole(c *fiber.Ctx) error {
	c.Locals(ADMIN_ROLE, enums.UserRole.ADMIN)
	return c.Next()
}

func UserRole(c *fiber.Ctx) error {
	c.Locals(USER_ROLE, enums.UserRole.USER)
	return c.Next()
}

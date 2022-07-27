package middlewares

import (
	"fiber-starter/errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func RoleBasedAuth(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err, ok := GetCurrentUser(c)
		if !ok {
			return err
		}
		if !slices.Contains(roles, user.Role) {
			return errors.ForbiddenException(c)
		}
		return c.Next()
	}
}

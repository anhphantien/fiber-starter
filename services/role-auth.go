package services

import (
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func RoleAuth(c *fiber.Ctx, user models.Claims, roles ...string) (error, bool) {
	if !slices.Contains(roles, user.Role) {
		return c.Status(fiber.StatusForbidden).JSON(HttpResponse{
			StatusCode: fiber.StatusForbidden,
			Error:      PERMISSION_DENIED,
		}), false
	}
	return nil, true
}

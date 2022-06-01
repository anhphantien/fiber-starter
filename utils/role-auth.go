package utils

import (
	"fiber-starter/common"
	"fiber-starter/errors"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func RoleAuth(c *fiber.Ctx, user models.JwtClaims, roles []string) (error, bool) {
	if !slices.Contains(roles, user.Role) {
		return c.Status(fiber.StatusForbidden).JSON(common.HttpResponse{
			StatusCode: fiber.StatusForbidden,
			Error:      errors.PERMISSION_DENIED,
		}), false
	}
	return nil, true
}

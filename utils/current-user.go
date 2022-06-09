package utils

import (
	"fiber-starter/errors"
	"fiber-starter/middlewares"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

func CurrentUser(c *fiber.Ctx) (models.JwtClaims, error, bool) {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := models.JwtClaims{
		ID:       uint64(claims["id"].(float64)),
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
	}
	if ok := validateUserRole(c, user); !ok {
		return user, errors.ForbiddenException(c), false
	}
	return user, nil, true
}

func validateUserRole(c *fiber.Ctx, user models.JwtClaims) bool {
	roles := []string{}

	if ADMIN, ok := c.Locals(middlewares.ADMIN_ROLE).(string); ok {
		roles = append(roles, ADMIN)
	}
	if USER, ok := c.Locals(middlewares.USER_ROLE).(string); ok {
		roles = append(roles, USER)
	}

	if len(roles) > 0 && !slices.Contains(roles, user.Role) {
		return false
	}
	return true
}

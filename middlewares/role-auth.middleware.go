package middlewares

import (
	"fiber-starter/enums"
	"fiber-starter/errors"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

const (
	ADMIN_ROLE = "ADMIN_ROLE"
	USER_ROLE  = "USER_ROLE"
)

func AdminRole(c *fiber.Ctx) error {
	c.Locals(ADMIN_ROLE, enums.User.Role.ADMIN)
	return c.Next()
}

func UserRole(c *fiber.Ctx) error {
	c.Locals(USER_ROLE, enums.User.Role.USER)
	return c.Next()
}

func RoleAuth(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := GetCurrentUser(c)
		if len(roles) > 0 && !slices.Contains(roles, user.Role) {
			return errors.ForbiddenException(c)
		}
		return c.Next()
	}
}

func GetCurrentUser(c *fiber.Ctx) models.CurrentUser {
	if c.Locals("user") == nil {
		return errors.UnauthorizedException(c)
	}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return user
}

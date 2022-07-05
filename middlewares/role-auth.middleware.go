package middlewares

import (
	"fiber-starter/errors"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

func RoleAuth(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := GetCurrentUser(c)
		if !ok {
			return errors.UnauthorizedException(c)
		}
		if len(roles) > 0 && !slices.Contains(roles, user.Role) {
			return errors.ForbiddenException(c)
		}
		return c.Next()
	}
}

func GetCurrentUser(c *fiber.Ctx) (models.CurrentUser, bool) {
	user := models.CurrentUser{}

	if c.Locals("user") == nil {
		return user, false
	}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return user, true
}

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
		user, err, ok := GetCurrentUser(c)
		if !ok {
			return err
		}
		if len(roles) > 0 && !slices.Contains(roles, user.Role) {
			return errors.ForbiddenException(c)
		}
		return c.Next()
	}
}

func GetCurrentUser(c *fiber.Ctx) (models.CurrentUser, error, bool) {
	currentUser := models.CurrentUser{}

	if c.Locals("user") == nil {
		return currentUser, errors.UnauthorizedException(c), false
	}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	currentUser = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return currentUser, nil, true
}

package middlewares

import (
	"fiber-starter/env"
	"fiber-starter/errors"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: env.JWT_SECRET,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			switch err.Error() {
			case "Missing or malformed JWT":
				return errors.UnauthorizedException(c, jwt.ErrTokenMalformed.Error())
			case "Invalid or expired JWT":
				return errors.UnauthorizedException(c, jwt.ErrTokenExpired.Error())
			default:
				return errors.UnauthorizedException(c, err.Error())
			}
		},
	})
}

func GetCurrentUser(c *fiber.Ctx) (models.CurrentUser, error, bool) {
	currentUser := models.CurrentUser{}

	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return currentUser, errors.InternalServerErrorException(c, errors.MISSING_JWT_AUTH), false
	}

	claims := user.Claims.(jwt.MapClaims)
	currentUser = models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}
	return currentUser, nil, true
}

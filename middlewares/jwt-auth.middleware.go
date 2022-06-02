package middlewares

import (
	"fiber-starter/env"
	"fiber-starter/errors"

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

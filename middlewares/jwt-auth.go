package middlewares

import (
	"fiber-starter/env"
	"fiber-starter/handlers"

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
				return c.Status(fiber.StatusUnauthorized).
					JSON(handlers.HttpResponse{
						StatusCode: fiber.StatusUnauthorized,
						Error:      jwt.ErrTokenMalformed.Error(),
					})
			case "Invalid or expired JWT":
				return c.Status(fiber.StatusUnauthorized).
					JSON(handlers.HttpResponse{
						StatusCode: fiber.StatusUnauthorized,
						Error:      jwt.ErrTokenExpired.Error(),
					})
			default:
				return c.Status(fiber.StatusUnauthorized).
					JSON(handlers.HttpResponse{
						StatusCode: fiber.StatusUnauthorized,
						Error:      err.Error(),
					})
			}
		},
	})
}

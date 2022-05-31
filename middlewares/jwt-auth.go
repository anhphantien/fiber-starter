package middlewares

import (
	"fiber-starter/env"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: jwtware.HS256,
		SigningKey:    env.JWT_SECRET,
	})
}

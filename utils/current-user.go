package utils

import (
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CurrentUser(c *fiber.Ctx) models.JwtClaims {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return models.JwtClaims{
		ID:       uint64(claims["id"].(float64)),
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
	}
}

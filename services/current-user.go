package services

import (
	"encoding/json"
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CurrentUser(c *fiber.Ctx) models.Claims {
	claims := c.Locals("user").(*jwt.Token).Claims
	jsonClaim, _ := json.Marshal(claims)
	user := models.Claims{}
	json.Unmarshal(jsonClaim, &user)
	return user
}

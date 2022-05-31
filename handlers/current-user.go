package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CurrentUser(c *fiber.Ctx) Claims {
	claims := c.Locals("user").(*jwt.Token).Claims
	jsonClaim, _ := json.Marshal(claims)
	user := Claims{}
	json.Unmarshal(jsonClaim, &user)
	return user
}

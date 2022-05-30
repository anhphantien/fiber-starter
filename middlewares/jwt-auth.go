package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtAuth(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), "Bearer ")[0]

	fmt.Println(token)
	// jwt.ParseWithClaims()

	return c.Next()
}

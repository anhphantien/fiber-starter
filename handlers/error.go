package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SqlError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(HttpResponse{
			StatusCode: fiber.StatusNotFound,
			Error:      err.Error(),
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(HttpResponse{
			StatusCode: fiber.StatusInternalServerError,
			Error:      err.Error(),
		})
	}
}

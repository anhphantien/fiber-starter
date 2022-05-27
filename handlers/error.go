package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DATA_NOT_FOUND   string
	INVALID_PASSWORD string
)

func init() {
	DATA_NOT_FOUND = "Data not found"
	INVALID_PASSWORD = "Invalid password"
}

func SqlError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(HttpResponse{
			StatusCode: fiber.StatusNotFound,
			Error:      DATA_NOT_FOUND,
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(HttpResponse{
			StatusCode: fiber.StatusInternalServerError,
			Error:      err.Error(),
		})
	}
}

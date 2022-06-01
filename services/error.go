package services

import (
	"fiber-starter/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DATA_NOT_FOUND    string
	INVALID_PASSWORD  string
	PERMISSION_DENIED string
)

func init() {
	DATA_NOT_FOUND = "data not found"
	INVALID_PASSWORD = "invalid password"
	PERMISSION_DENIED = "permission denied"
}

func SqlError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(models.HttpResponse{
			StatusCode: fiber.StatusNotFound,
			Error:      DATA_NOT_FOUND,
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(models.HttpResponse{
			StatusCode: fiber.StatusInternalServerError,
			Error:      err.Error(),
		})
	}
}

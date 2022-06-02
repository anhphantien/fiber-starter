package errors

import (
	"fiber-starter/common"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DATA_NOT_FOUND    = "data not found"
	INVALID_PASSWORD  = "invalid password"
	PERMISSION_DENIED = "permission denied"
)

func BadRequestException(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
		StatusCode: fiber.StatusBadRequest,
		Message:    message,
	})
}

func UnauthorizedException(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(common.HttpResponse{
		StatusCode: fiber.StatusUnauthorized,
		Message:    message,
	})
}

func ForbiddenException(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message[0] = PERMISSION_DENIED
	}
	return c.Status(fiber.StatusForbidden).JSON(common.HttpResponse{
		StatusCode: fiber.StatusForbidden,
		Message:    message[0],
	})
}

func NotFoundException(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message[0] = DATA_NOT_FOUND
	}
	return c.Status(fiber.StatusNotFound).JSON(common.HttpResponse{
		StatusCode: fiber.StatusNotFound,
		Message:    message[0],
	})
}

func InternalServerErrorException(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(common.HttpResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    message,
	})
}

func SqlError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return NotFoundException(c)
	default:
		return InternalServerErrorException(c, err.Error())
	}
}

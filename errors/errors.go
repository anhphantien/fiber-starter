package errors

import (
	"fiber-starter/common"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	DATA_NOT_FOUND      = "data not found"
	FILE_NOT_FOUND      = "file not found"
	INVALID_FILE_FORMAT = "invalid file format"
	INVALID_PASSWORD    = "invalid password"
	PAYLOAD_TOO_LARGE   = "payload too large"
	PERMISSION_DENIED   = "permission denied"
)

func BadRequestException(c *fiber.Ctx, message string) error {
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusBadRequest,
		Message:    message,
	})
}

func UnauthorizedException(c *fiber.Ctx, message string) error {
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusUnauthorized,
		Message:    message,
	})
}

func ForbiddenException(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		return common.HttpResponse(c, common.Response{
			StatusCode: fiber.StatusForbidden,
			Message:    PERMISSION_DENIED,
		})
	}
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusForbidden,
		Message:    message[0],
	})
}

func NotFoundException(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		return common.HttpResponse(c, common.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    DATA_NOT_FOUND,
		})
	}
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusNotFound,
		Message:    message[0],
	})
}

func RequestEntityTooLargeException(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		return common.HttpResponse(c, common.Response{
			StatusCode: fiber.StatusRequestEntityTooLarge,
			Message:    PAYLOAD_TOO_LARGE,
		})
	}
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusRequestEntityTooLarge,
		Message:    message[0],
	})
}

func InternalServerErrorException(c *fiber.Ctx, message string) error {
	return common.HttpResponse(c, common.Response{
		StatusCode: fiber.StatusInternalServerError,
		Message:    message,
	})
}

func SqlError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return NotFoundException(c, err.Error())
	default:
		return InternalServerErrorException(c, err.Error())
	}
}

package errors

import (
	"fiber-starter/response"

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

func BadRequestException(c *fiber.Ctx, err any) error {
	switch err := err.(type) {
	case string:
		return response.WriteJSON(c, response.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    err,
		})
	default:
		return response.WriteJSON(c, response.Response{
			StatusCode: fiber.StatusBadRequest,
			Error:      err.([]response.Error),
		})
	}
}

func UnauthorizedException(c *fiber.Ctx, message string) error {
	return response.WriteJSON(c, response.Response{
		StatusCode: fiber.StatusUnauthorized,
		Message:    message,
	})
}

func ForbiddenException(c *fiber.Ctx, messages ...string) error {
	message := PERMISSION_DENIED
	if len(messages) > 0 {
		message = messages[0]
	}
	return response.WriteJSON(c, response.Response{
		StatusCode: fiber.StatusForbidden,
		Message:    message,
	})
}

func NotFoundException(c *fiber.Ctx, messages ...string) error {
	message := DATA_NOT_FOUND
	if len(messages) > 0 {
		message = messages[0]
	}
	return response.WriteJSON(c, response.Response{
		StatusCode: fiber.StatusNotFound,
		Message:    message,
	})
}

func RequestEntityTooLargeException(c *fiber.Ctx, messages ...string) error {
	message := PAYLOAD_TOO_LARGE
	if len(messages) > 0 {
		message = messages[0]
	}
	return response.WriteJSON(c, response.Response{
		StatusCode: fiber.StatusRequestEntityTooLarge,
		Message:    message,
	})
}

func InternalServerErrorException(c *fiber.Ctx, message string) error {
	return response.WriteJSON(c, response.Response{
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

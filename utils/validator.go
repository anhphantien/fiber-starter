package utils

import (
	"fiber-starter/errors"
	"fiber-starter/response"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequestBody(c *fiber.Ctx, payload any) (error, bool) {
	if err := c.BodyParser(payload); err != nil {
		return errors.BadRequestException(c, err.Error()), false
	}

	if err := validator.New().Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		err := make([]response.Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			err[i] = response.Error{
				Field: strings.ToLower(fieldError.Field()),
				Message: func(fieldError validator.FieldError) string {
					switch fieldError.Tag() {
					case "required":
						return "This field is required"
					case "max":
						return "Max length: " + fieldError.Param()
					default:
						return fieldError.Error()
					}
				}(fieldError),
			}
		}
		return errors.BadRequestException(c, err), false
	}
	return nil, true
}

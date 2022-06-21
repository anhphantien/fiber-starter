package utils

import (
	"fiber-starter/common"
	"fiber-starter/errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequestBody(c *fiber.Ctx, payload any) (error, bool) {
	validate := validator.New()

	if err := c.BodyParser(payload); err != nil {
		return errors.BadRequestException(c, err.Error()), false
	}

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			errors[i] = Error{
				Field:   strings.ToLower(fieldError.Field()),
				Message: message(fieldError),
			}
		}

		return common.HttpResponse(c, common.Response{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors,
		}), false
	}

	return nil, true
}

func message(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return "Max length: " + fieldError.Param()
	}
	return fieldError.Error()
}

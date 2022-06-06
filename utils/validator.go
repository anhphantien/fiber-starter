package utils

import (
	"fiber-starter/common"
	"fiber-starter/errors"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validate(c *fiber.Ctx, payload any) (error, bool) {
	validate := validator.New()

	if err := c.BodyParser(payload); err != nil {
		return errors.BadRequestException(c, err.Error()), false
	}

	if err := validate.Struct(payload); err != nil {
		_error := []ApiError{}

		for _, fieldError := range err.(validator.ValidationErrors) {
			_error = append(_error, ApiError{
				Field:   makeFirstLetterLowercase(fieldError.Field()),
				Message: message(fieldError),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(common.HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      _error,
		}), false
	}

	return nil, true
}

func makeFirstLetterLowercase(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
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

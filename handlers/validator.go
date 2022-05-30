package handlers

import (
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validate(c *fiber.Ctx, payload interface{}) (error, bool) {
	validate := validator.New()

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Error:      err.Error(),
		}), false
	}

	if err := validate.Struct(payload); err != nil {
		errors := []ApiError{}

		for _, fieldError := range err.(validator.ValidationErrors) {
			errors = append(errors, ApiError{Field: makeFirstLetterLowercase(fieldError.Field()), Message: msgForTag(fieldError)})
		}

		return c.Status(fiber.StatusBadRequest).JSON(HttpResponse{
			StatusCode: fiber.StatusBadRequest,
			Errors:     errors,
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

func msgForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	}
	return fieldError.Error()
}

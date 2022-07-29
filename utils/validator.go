package utils

import (
	"fiber-starter/errors"
	"fiber-starter/response"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequestBody(c *fiber.Ctx, payload any) (error, bool) {
	c.BodyParser(payload)

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		err := make([]response.Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			err[i] = response.Error{
				Field: strings.ToLower(fieldError.Field()),
				Message: func(fe validator.FieldError) string {
					switch fe.Tag() {
					case "required":
						return "This field is required"
					case "max":
						return "Max length: " + fe.Param()
					default:
						return fe.Error()
					}
				}(fieldError),
			}
		}
		return errors.BadRequestException(c, err), false
	}
	return nil, true
}

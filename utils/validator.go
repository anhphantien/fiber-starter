package utils

import (
	"encoding/json"
	"fiber-starter/errors"
	"fiber-starter/response"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

func ValidateRequestBody(c *fiber.Ctx, body any) (error, bool) {
	if err := c.BodyParser(body); err != nil {
		return errors.BadRequestException(c, err.Error()), false
	}

	if err := validator.New().Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		_errors := make([]response.Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			_errors[i] = response.Error{
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

		return errors.BadRequestException(c, _errors), false
	}

	return nil, true
}

func FilterRequestBody(c *fiber.Ctx, body any) map[string]any {
	dto := map[string]any{}
	_dto, _ := json.Marshal(body)
	json.Unmarshal(_dto, &dto)

	rawBody := map[string]any{}
	json.Unmarshal(c.Body(), &rawBody)

	return lo.PickByKeys(rawBody, maps.Keys(dto))
}

package utils

import (
	"encoding/json"
	"fiber-starter/common"
	"fiber-starter/errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequestBody(c *fiber.Ctx, payload any) (map[string]any, error, bool) {
	validate := validator.New()

	if err := c.BodyParser(payload); err != nil {
		return nil, errors.BadRequestException(c, err.Error()), false
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

		return nil, common.HttpResponse(c, common.Response{
			StatusCode: fiber.StatusBadRequest,
			Error:      errors,
		}), false
	}

	dto := map[string]any{}
	_payload, _ := json.Marshal(payload)
	json.Unmarshal(_payload, &dto)

	body := map[string]any{}
	json.Unmarshal(c.Body(), &body)

	data := lo.PickByKeys(body, maps.Keys(dto))

	return data, nil, true
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

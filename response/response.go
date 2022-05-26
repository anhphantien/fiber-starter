package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Http struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func Error(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(http.StatusNotFound).JSON(Http{
			StatusCode: http.StatusNotFound,
			Error:      err.Error(),
		})
	default:
		return c.Status(http.StatusInternalServerError).JSON(Http{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
	}
}

func Validate(c *fiber.Ctx, payload interface{}) error {
	var validate = validator.New()

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Http{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	if err := validate.Struct(payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Http{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	return nil
}

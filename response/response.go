package response

import (
	"net/http"

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

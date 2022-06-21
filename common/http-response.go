package common

import "github.com/gofiber/fiber/v2"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	Error      any    `json:"error,omitempty"`
}

func HttpResponse(c *fiber.Ctx, response Response) error {
	if response.StatusCode == 0 && c.Route().Method == fiber.MethodPost {
		response.StatusCode = fiber.StatusCreated
	}
	return c.Status(response.StatusCode).JSON(response)
}

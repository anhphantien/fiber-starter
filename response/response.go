package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	Error      any    `json:"error,omitempty"`
}

func WriteJSON(c *fiber.Ctx, payload Response) error {
	if payload.StatusCode == 0 {
		if c.Route().Method == fiber.MethodPost {
			payload.StatusCode = fiber.StatusCreated
		} else {
			payload.StatusCode = fiber.StatusOK
		}
	}
	return c.Status(payload.StatusCode).JSON(payload)
}

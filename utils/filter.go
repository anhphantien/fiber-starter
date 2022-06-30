package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

func FilterRequestBody(c *fiber.Ctx, payload any) map[string]any {
	dto := map[string]any{}
	_dto, _ := json.Marshal(payload)
	json.Unmarshal(_dto, &dto)

	rawBody := map[string]any{}
	json.Unmarshal(c.Body(), &rawBody)

	return lo.PickByKeys(rawBody, maps.Keys(dto))
}

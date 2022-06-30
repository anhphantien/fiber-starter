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

	body := map[string]any{}
	json.Unmarshal(c.Body(), &body)

	return lo.PickByKeys(body, maps.Keys(dto))
}

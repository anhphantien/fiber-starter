package utils

import (
	"encoding/json"
	"fiber-starter/dto"
	"fiber-starter/enums"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func Pagination(c *fiber.Ctx) dto.Pagination {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}

	keyword := c.Query("keyword")

	filter := map[string]any{}
	json.Unmarshal([]byte(c.Query("filter")), &filter)

	sort := dto.Sort{}
	json.Unmarshal([]byte(c.Query("sort")), &sort)
	if len(sort.Field) == 0 {
		sort.Field = "id"
	}
	if !slices.Contains([]string{enums.SortOrder.ASC, enums.SortOrder.DESC}, sort.Order) {
		sort.Order = enums.SortOrder.DESC
	}

	return dto.Pagination{
		Limit:   limit,
		Page:    page,
		Keyword: keyword,
		Filter:  filter,
		Sort:    sort,
	}
}

package utils

import (
	"encoding/json"
	"fiber-starter/dto"
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

	var sort struct {
		By        string
		Direction string
	}
	json.Unmarshal([]byte(c.Query("sort")), &sort)
	if len(sort.By) == 0 {
		sort.By = "id"
	}
	if !slices.Contains(
		[]string{
			"ASC",
			"DESC",
		}, sort.Direction) {
		sort.Direction = "DESC"
	}

	return dto.Pagination{
		Limit:   limit,
		Offset:  limit * (page - 1),
		Keyword: keyword,
		Filter:  filter,
		Order:   sort.By + " " + sort.Direction,
	}
}

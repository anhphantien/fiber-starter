package utils

import (
	"encoding/json"
	"fiber-starter/dto"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Pagination(c *fiber.Ctx) dto.Pagination {
	fmt.Println(c.Query("limit"))
	fmt.Println(c.Query("page"))
	fmt.Println(c.Query("keyword"))
	fmt.Println(c.Query("filter"))
	fmt.Println(c.Query("sort"))

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

	fmt.Println(11111111111, sort.Field)
	fmt.Println(22222222222, sort.Order)

	return dto.Pagination{
		Limit:   limit,
		Page:    page,
		Keyword: keyword,
		Filter:  filter,
		Sort:    sort,
	}
}

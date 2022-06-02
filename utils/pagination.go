package utils

import (
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
	if limit < 1 {
		limit = 1
	}

	keyword := c.Query("keyword")

	return dto.Pagination{
		Limit:   limit,
		Page:    page,
		Keyword: keyword,
	}
}

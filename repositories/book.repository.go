package repositories

import (
	"fiber-starter/dto"
	"fiber-starter/errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookRepository struct{}

func (r BookRepository) GetCount(c *fiber.Ctx, q *gorm.DB, total int64) <-chan error {
	ch := make(chan error)

	go q.Count(&total)
	if q.Error != nil {
		ch <- errors.SqlError(c, q.Error)
	}
	ch <- nil

	return ch
}

func (r BookRepository) GetItems(c *fiber.Ctx, q *gorm.DB, pagination dto.Pagination, items any) <-chan error {
	ch := make(chan error)

	go q.Limit(pagination.Limit).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Order(pagination.Sort.Field + " " + pagination.Sort.Order).
		Find(&items)
	if q.Error != nil {
		ch <- errors.SqlError(c, q.Error)
	}
	ch <- nil

	return ch
}

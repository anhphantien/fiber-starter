package services

import (
	"fiber-starter/common"
	"fiber-starter/database"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct{}

// @Tags    users
// @Summary Get a list of users
// @Param   limit     query    int false " "
// @Param   page      query    int false " "
// @Param   keyword   query    string false " "
// @Param   filter    query    object false " "
// @Param   sort      query    object false " "
// @Success 200       {object} common.HttpResponse{data=[]entities.User}
// @Router  /v1/users [get]
func (s UserService) GetList(c *fiber.Ctx) error {
	db := database.DB

	users := []entities.User{}

	pagination := utils.Pagination(c)

	q := db.
		Model(users).
		Preload("Books", "TRUE ORDER BY book.id DESC")

	ch := make(chan error, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var total int64
	go func() {
		defer wg.Done()

		r := q.
			Session(&gorm.Session{}). // clone
			Count(&total)
		if r.Error != nil {
			ch <- r.Error
		}
	}()

	go func() {
		defer wg.Done()

		r := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Limit * (pagination.Page - 1)).
			Order(pagination.Sort.Field + " " + pagination.Sort.Order).
			Find(&users)
		if r.Error != nil {
			ch <- r.Error
		}
	}()

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return errors.SqlError(c, err)
		}
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data: common.PaginationResponse{
			Items: users,
			Total: total,
		},
	})
}

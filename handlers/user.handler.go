package handlers

import (
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/models"
	"fiber-starter/repositories"
	"fiber-starter/response"
	"fiber-starter/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct{}

// @Tags    users
// @Summary Get a list of users
// @Param   limit         query    int    false " "
// @Param   page          query    int    false " "
// @Param   keyword       query    string false " "
// @Param   filter        query    object false " "
// @Param   sort          query    object false " "
// @Success 200           {object} response.Response{data=[]entities.User}
// @Router  /api/v1/users [GET]
func (h UserHandler) GetList(c *fiber.Ctx) error {
	users := []entities.User{}

	pagination := utils.Pagination(c)

	q := repositories.CreateSqlBuilder(users).
		Preload("Books", "TRUE ORDER BY id DESC")

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
			Offset(pagination.Offset).
			Order(pagination.Order).
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

	return response.WriteJSON(c, response.Response{
		Data: models.PaginationResponse{
			Items: users,
			Total: total,
		},
	})
}

package handlers

import (
	"database/sql"
	"fiber-starter/dto"
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

type BookHandler struct{}

// @Tags    books
// @Summary Get a list of books
// @Param   limit         query  int    false " "
// @Param   page          query  int    false " "
// @Param   keyword       query  string false " "
// @Param   filter        query  object false " "
// @Param   sort          query  object false " "
// @Success 200           object response.Response{data=[]entities.Book}
// @Router  /api/v1/books [GET]
func (h BookHandler) GetList(c *fiber.Ctx) error {
	books := []entities.Book{}

	pagination := utils.Pagination(c)

	q := repositories.CreateSqlBuilder(books).
		Preload("User")
	if pagination.Filter["id"] != nil {
		q.Where("book.id = ?", utils.ConvertToID(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"book.title LIKE @keyword OR book.description LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}

	// var err error

	// var total int64
	// err = q.Count(&total).Error
	// if err != nil {
	// 	return errors.SqlError(c, err)
	// }

	// err = q.
	// 	Limit(pagination.Limit).
	// 	Offset(pagination.Offset).
	// 	Order(pagination.Order).
	// 	Find(&books).Error
	// if err != nil {
	// 	return errors.SqlError(c, err)
	// }

	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	var total int64
	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Count(&total).Error
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Order(pagination.Order).
			Find(&books).Error
		if err != nil {
			ch <- err
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
			Items: books,
			Total: total,
		},
	})
}

// @Tags    books
// @Summary Get a book by ID
// @Param   id                 path   int true " "
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [GET]
func (h BookHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err, ok := bookRepository.FindOneByID(c, id)
	if !ok {
		return err
	}

	return response.WriteJSON(c, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Create a new book
// @Param   body          body   dto.CreateBookBody true " "
// @Success 201           object response.Response{data=entities.Book}
// @Router  /api/v1/books [POST]
func (h BookHandler) Create(c *fiber.Ctx) error {
	body := dto.CreateBookBody{}
	if err, ok := utils.ValidateRequestBody(c, &body); !ok {
		return err
	}

	if body.UserID != nil {
		_, err, ok := userRepository.FindOneByID(c, body.UserID)
		if !ok {
			return err
		}
	}

	book, err, ok := bookRepository.Create(c, body)
	if !ok {
		return err
	}

	return response.WriteJSON(c, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Update a book
// @Param   id                 path   int true " "
// @Param   body               body   dto.UpdateBookBody true " "
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [PUT]
func (h BookHandler) Update(c *fiber.Ctx) error {
	body := dto.UpdateBookBody{}
	if err, ok := utils.ValidateRequestBody(c, &body); !ok {
		return err
	}

	book, err, ok := bookRepository.Update(c, body)
	if !ok {
		return err
	}

	return response.WriteJSON(c, response.Response{
		Data: book,
	})
}

// @Security Bearer
// @Summary  Delete a book
// @Tags     books
// @Param    id                 path     int true " "
// @Success  200                object   response.Response{data=boolean}
// @Router   /api/v1/books/{id} [DELETE]
func (h BookHandler) Delete(c *fiber.Ctx) error {
	// user, err, ok := middlewares.GetCurrentUser(c)
	// if !ok {
	// 	return err
	// }

	id := c.Params("id")

	err, ok := bookRepository.Delete(c, id)
	if !ok {
		return err
	}

	return response.WriteJSON(c, response.Response{
		Data: true,
	})
}

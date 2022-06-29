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

	// var total int64
	// r1 := q.Count(&total)
	// if r1.Error != nil {
	// 	return errors.SqlError(c, r1.Error)
	// }

	// r2 := q.
	// 	Limit(pagination.Limit).
	// 	Offset(pagination.Offset).
	// 	Order(pagination.Order).
	// 	Find(&books)
	// if r2.Error != nil {
	// 	return errors.SqlError(c, r2.Error)
	// }

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
			Find(&books)
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

	book, err := repositories.BookRepository{}.FindOneByID(id)
	if err != nil {
		return errors.SqlError(c, err)
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
		_, err := repositories.UserRepository{}.FindOneByID(body.UserID)
		if err != nil {
			return errors.SqlError(c, err)
		}
	}

	book, err := repositories.BookRepository{}.Create(body)
	if err != nil {
		return errors.SqlError(c, err)
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

	book, err := repositories.BookRepository{}.Update(c, body)
	if err != nil {
		return errors.SqlError(c, err)
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
	// user, err, ok := utils.GetCurrentUser(c)
	// if !ok {
	// 	return err
	// }

	id := c.Params("id")

	err := repositories.BookRepository{}.Delete(id)
	if err != nil {
		return errors.SqlError(c, err)
	}

	return response.WriteJSON(c, response.Response{})
}

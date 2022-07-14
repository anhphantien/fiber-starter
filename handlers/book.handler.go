package handlers

import (
	"database/sql"
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/models"
	"fiber-starter/repositories"
	"fiber-starter/response"
	"fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
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
	pagination := utils.Pagination(c)

	books := []entities.Book{}
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
	q.Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(pagination.Order)
	books, total, err, ok := bookRepository.FindAndCount(c, q)
	if !ok {
		return err
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

	book, err, ok := bookRepository.FindByID(c, id)
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
		_, err, ok := userRepository.FindByID(c, body.UserID)
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
	// currentUser, err, ok := middlewares.GetCurrentUser(c)
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

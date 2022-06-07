package services

import (
	"database/sql"
	"fiber-starter/common"
	"fiber-starter/database"
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BookService struct{}

// @Summary Get all books
// @Tags books
// @Param limit query int false " "
// @Param page query int false " "
// @Param keyword query string false " "
// @Param filter query object false " "
// @Param sort query object false " "
// @Success 200 {object} common.HttpResponse{data=[]entities.Book}
// @Router /v1/books [get]
func (h BookService) GetList(c *fiber.Ctx) error {
	db := database.DB

	books := []entities.Book{}

	pagination := utils.Pagination(c)

	q := db.Model(books)
	if pagination.Filter["id"] != nil {
		q.Where("id = ?", utils.ConvertToInt(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"title LIKE @keyword OR description LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}

	// var total int64
	// q.Count(&total)

	// q.Limit(pagination.Limit).
	// 	Offset(pagination.Limit * (pagination.Page - 1)).
	// 	Order(pagination.Sort.Field + " " + pagination.Sort.Order).
	// 	Find(&books) // db.Table("book").Select("1 + 2 AS sum, \"abc\" AS title").Scan(&books)
	// if q.Error != nil {
	// 	return errors.SqlError(c, q.Error)
	// }

	wg := sync.WaitGroup{}
	wg.Add(2)

	var total int64
	go func() error {
		defer wg.Done()

		r1 := q.
			Session(&gorm.Session{}). // clone
			Count(&total)
		if r1.Error != nil {
			return errors.SqlError(c, q.Error)
		}
		return nil
	}()

	go func() error {
		defer wg.Done()

		r2 := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Limit * (pagination.Page - 1)).
			Order(pagination.Sort.Field + " " + pagination.Sort.Order).
			Find(&books)
		if r2.Error != nil {
			return errors.SqlError(c, q.Error)
		}
		return nil
	}()
	wg.Wait()

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data: common.PaginationResponse{
			Items: books,
			Total: total,
		},
	})
}

// @Summary Get a book by ID
// @Tags books
// @Param id path int true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books/{id} [get]
func (h BookService) GetByID(c *fiber.Ctx) error {
	db := database.DB

	book := entities.Book{}
	id := utils.ConvertToInt(c.Params("id"))

	r := db.Model(book).Where("id = ?", id).First(&book)
	if r.Error != nil {
		return errors.SqlError(c, r.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
		Data:       book,
	})
}

// @Summary Create a new book
// @Tags books
// @Param body body dto.CreateBookBody true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books [post]
func (h BookService) Create(c *fiber.Ctx) error {
	db := database.DB

	body := dto.CreateBookBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	book := entities.Book{}
	copier.Copy(&book, &body)

	r := db.Create(&book)
	if r.Error != nil {
		return errors.SqlError(c, r.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Summary Update a book
// @Tags books
// @Param id path int true " "
// @Param body body dto.UpdateBookBody true " "
// @Success 200 {object} common.HttpResponse{data=entities.Book}
// @Router /v1/books/{id} [put]
func (h BookService) Update(c *fiber.Ctx) error {
	db := database.DB

	body := dto.UpdateBookBody{}
	if err, ok := utils.Validate(c, &body); !ok {
		return err
	}

	book := entities.Book{}
	id := utils.ConvertToInt(c.Params("id"))

	q := db.Model(book).Session(&gorm.Session{})
	r1 := q.Where("id = ?", id).First(&book)
	if r1.Error != nil {
		return errors.SqlError(c, r1.Error)
	}

	copier.Copy(&book, &body)
	r2 := q.Where("id = ?", id).Updates(&book)
	if r2.Error != nil {
		return errors.SqlError(c, r2.Error)
	}

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
		Data:       book,
	})
}

// @Security BearerAuth
// @Summary Delete a book
// @Tags books
// @Param id path int true " "
// @Success 200 {object} common.HttpResponse{}
// @Router /v1/books/{id} [delete]
func (h BookService) Delete(c *fiber.Ctx) error {
	db := database.DB

	// user, err, ok := utils.CurrentUser(c)
	// if !ok {
	// 	return err
	// }

	book := entities.Book{}

	if err := db.First(&book, c.Params("id")).Error; err != nil {
		return errors.SqlError(c, err)
	}

	db.Delete(&book)

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusOK,
	})
}

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

	q := db.Model(&books)
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
	go func() {
		defer wg.Done()

		q.
			Session(&gorm.Session{}). // clone
			Count(&total)
	}()

	go func() error {
		defer wg.Done()

		q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Limit * (pagination.Page - 1)).
			Order(pagination.Sort.Field + " " + pagination.Sort.Order).
			Find(&books)
		if q.Error != nil {
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
	q := db.Model(&book).Where("id = ?", id).First(&book)
	if q.Error != nil {
		return errors.SqlError(c, q.Error)
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

	q := db.Create(&book)
	if q.Error != nil {
		return errors.SqlError(c, q.Error)
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
	q := db.Model(&book).Where("id = ?", id).First(&book)
	if q.Error != nil {
		return errors.SqlError(c, q.Error)
	}

	copier.Copy(&book, &body)
	q.Updates(&book)

	// q := db.Model(&book).Session(&gorm.Session{DryRun: true})
	// q.
	// 	Where("id = ?", id).
	// 	First(&book)
	// if q.Error != nil {
	// 	return errors.SqlError(c, q.Error)
	// }

	// copier.Copy(&book, &body)
	// q.
	// 	Where("id = ?", 2).
	// 	Updates(&book)
	// if q.Error != nil {
	// 	return errors.SqlError(c, q.Error)
	// }

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

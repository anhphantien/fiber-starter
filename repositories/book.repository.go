package repositories

import (
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
)

var book = entities.Book{}

type BookRepository struct{}

func (repository BookRepository) FindOneByID(c *fiber.Ctx, id any) (entities.Book, error, bool) {
	err := CreateSqlBuilder(book).
		Joins("User").
		Where("book.id = ?", utils.ConvertToID(id)).
		Take(&book).Error
	if err != nil {
		return book, errors.SqlError(c, err), false
	}
	return book, nil, true
}

func (repository BookRepository) Create(c *fiber.Ctx, body dto.CreateBookBody) (entities.Book, error, bool) {
	copier.Copy(&book, body)
	err := CreateSqlBuilder(book).Create(&book).Error
	if err != nil {
		return book, errors.SqlError(c, err), false
	}
	return book, nil, true
}

func (repository BookRepository) Update(c *fiber.Ctx, body dto.UpdateBookBody) (entities.Book, error, bool) {
	id := c.Params("id")

	book, err, ok := repository.FindOneByID(c, id)
	if !ok {
		return book, errors.SqlError(c, err), false
	}

	copier.Copy(&book, body)
	err = CreateSqlBuilder(book).
		Omit(clause.Associations). // skip auto create/update
		Updates(utils.FilterRequestBody(c, body)).Error
	if err != nil {
		return book, errors.SqlError(c, err), false
	}
	return book, nil, true
}

func (repository BookRepository) Delete(c *fiber.Ctx, id any) (error, bool) {
	book, err, ok := repository.FindOneByID(c, id)
	if !ok {
		return errors.SqlError(c, err), false
	}

	err = CreateSqlBuilder(book).Delete(&book).Error
	if !ok {
		return errors.SqlError(c, err), false
	}
	return nil, true
}

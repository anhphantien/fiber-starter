package repositories

import (
	"fiber-starter/dto"
	"fiber-starter/entities"
	"fiber-starter/errors"
	"fiber-starter/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var book = entities.Book{}

type BookRepository struct{}

func (repository BookRepository) FindAndCount(c *fiber.Ctx, q *gorm.DB) ([]entities.Book, int64, error, bool) {
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

	var books = []entities.Book{}
	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Find(&books).Error
		if err != nil {
			ch <- err
		}
	}()

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return books, total, errors.SqlError(c, err), false
		}
	}
	return books, total, nil, true
}

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
		Omit(clause.Associations). // skip all associations
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

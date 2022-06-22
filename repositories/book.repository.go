package repositories

import (
	"fiber-starter/entities"
	"fiber-starter/utils"
)

type BookRepository struct{}

func (r BookRepository) FindOneByID(id any) (book entities.Book, err error) {
	err = CreateSqlBuilder(book).
		Where("id = ?", utils.ConvertToInt(id)).
		Take(&book).Error
	return book, err
}

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

func (r BookRepository) Create(book entities.Book) (err error) {
	err = DB.Create(book).Error
	return err
}

package repositories

import (
	"fiber-starter/entities"
	"fiber-starter/utils"

	"github.com/jinzhu/copier"
)

type BookRepository struct{}

func (r BookRepository) FindOneByID(id any) (book entities.Book, err error) {
	err = CreateSqlBuilder(book).
		Where("id = ?", utils.ConvertToID(id)).
		Take(&book).Error
	return book, err
}

func (r BookRepository) Create(body any) (book entities.Book, err error) {
	copier.Copy(&book, body)
	err = DB.Create(&book).Error
	return book, err
}

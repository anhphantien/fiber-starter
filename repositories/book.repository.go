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

func (r BookRepository) Create(data any) (book entities.Book, err error) {
	copier.Copy(&book, data)
	err = CreateSqlBuilder(book).Create(&book).Error
	return book, err
}

func (r BookRepository) Update(id, data any) (book entities.Book, err error) {
	book, err = BookRepository{}.FindOneByID(id)
	if err != nil {
		return book, err
	}
	err = CreateSqlBuilder(book).Updates(data).Error
	return book, err
}

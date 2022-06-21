package repositories

import (
	"fiber-starter/entities"

	"gorm.io/gorm"
)

var BookRepository *gorm.DB

func (r Repository) BookRepository(db *gorm.DB) {
	book := entities.Book{}
	db.AutoMigrate(book)
	BookRepository = db.Model(book)
}

package repositories

import (
	"fiber-starter/entities"

	"gorm.io/gorm"
)

var BookRepository *gorm.DB

func _BookRepository(db *gorm.DB) {
	db.AutoMigrate(entities.Book{})
	BookRepository = db.Model(entities.Book{})
}

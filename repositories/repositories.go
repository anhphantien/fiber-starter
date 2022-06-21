package repositories

import (
	"gorm.io/gorm"
)

type Repository struct{}

func Init(db *gorm.DB) {
	_BookRepository(db)
	_UserRepository(db)
}

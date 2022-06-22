package repositories

import (
	"gorm.io/gorm"
)

type Repository struct{}

func New(db *gorm.DB) {
	r := Repository{}
	r.BookRepository(db)
	r.UserRepository(db)
}

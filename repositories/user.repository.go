package repositories

import (
	"fiber-starter/entities"

	"gorm.io/gorm"
)

var UserRepository *gorm.DB

func (r Repository) UserRepository(db *gorm.DB) {
	user := entities.User{}
	db.AutoMigrate(user)
	UserRepository = db.Model(user)
}

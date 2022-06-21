package repositories

import (
	"fiber-starter/entities"

	"gorm.io/gorm"
)

var UserRepository *gorm.DB

func _UserRepository(db *gorm.DB) {
	db.AutoMigrate(entities.User{})
	UserRepository = db.Model(entities.User{})
}

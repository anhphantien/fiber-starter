package repositories

import (
	"fiber-starter/database"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return database.DB
}

func GetRepository(model any) *gorm.DB {
	return database.DB.Model(model)
}

package database

import (
	"fiber-starter/env"
	"fiber-starter/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() (err error) {
	var (
		user     = env.DB_USER
		password = env.DB_PASS
		host     = env.DB_HOST
		port     = env.DB_PORT
		dbname   = env.DB_NAME
	)

	dsn := fmt.Sprint(user, password, ":@tcp(", host, ":", port, ")/", dbname, "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	DB.AutoMigrate(&models.Book{}, &models.User{})

	return
}

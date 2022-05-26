package database

import (
	"fiber-starter/env"
	"fiber-starter/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func Connect() (err error) {
	var (
		user     = env.USER
		password = env.PASSWORD
		host     = env.HOST
		dbname   = env.DB_NAME
		port     = env.PORT
	)

	dsn := fmt.Sprint(user, password, ":@tcp(", host, ":", port, ")/", dbname, "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
	fmt.Println(dsn)

	DBConn, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	DBConn.Logger = logger.Default.LogMode(logger.Info)
	DBConn.AutoMigrate(&models.Book{})
	return
}

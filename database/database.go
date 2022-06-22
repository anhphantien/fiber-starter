package database

import (
	"fiber-starter/env"
	"fiber-starter/repositories"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	username = env.DB_USER
	password = env.DB_PASS
	host     = env.DB_HOST
	port     = env.DB_PORT
	dbname   = env.DB_NAME
)

func Connect() (err error) {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true"
	fmt.Println(dsn)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		return err
	}

	repositories.Sync(db)
	return
}

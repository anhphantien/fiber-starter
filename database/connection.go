package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn   *gorm.DB
	user     string
	password string
	host     string
	db       string
	port     string
)

func Connect() (err error) {
	godotenv.Load()

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	host = os.Getenv("DB_HOST")
	db = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")

	dsn := fmt.Sprint(user, password, ":@tcp(", host, ":", port, ")/", db, "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
	fmt.Println(dsn)

	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return
}

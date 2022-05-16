package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn   *gorm.DB
	user     = os.Getenv("DB_USER")
	password = ""
	host     = os.Getenv("DB_HOST")
	db       = os.Getenv("DB_NAME")
	port     = os.Getenv("DB_PORT")
)

// Connect creates a connection to database
func Connect() (err error) {
	// port, err := strconv.Atoi(port)
	// if err != nil {
	// 	return err
	// }

	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&collation=utf8mb4_unicode_ci"
	fmt.Println(dsn)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := DBConn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

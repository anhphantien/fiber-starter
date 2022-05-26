package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	USER     string
	PASSWORD string
	HOST     string
	DB_NAME  string
	PORT     string
)

func init() {
	godotenv.Load(".env")

	USER = os.Getenv("DB_USER")
	PASSWORD = os.Getenv("DB_PASS")
	HOST = os.Getenv("DB_HOST")
	PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
}

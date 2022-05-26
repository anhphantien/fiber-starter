package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string

	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
)

func init() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
}

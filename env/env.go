package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	User     string
	Password string
	Host     string
	Db       string
	Port     string
)

func init() {
	godotenv.Load(".env")

	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASS")
	Host = os.Getenv("DB_HOST")
	Db = os.Getenv("DB_NAME")
	Port = os.Getenv("DB_PORT")
}

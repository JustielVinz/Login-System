package middleware

import (
	"log"
	"os"
	database "sample/middleware/utils/database"

	"github.com/joho/godotenv"
)

func CreateConnection() {
	database.PostgreSQLConnect(
		GetEnv("POSTGRES_USERNAME"),
		GetEnv("POSTGRES_PASSWORD"),
		GetEnv("POSTGRES_HOST"),
		GetEnv("DATABASE_NAME"),
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_SSL_MODE"),
		GetEnv("POSTGRES_TIMEZONE"),
	)

}

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

package setupDB

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ClientDB() (*gorm.DB, string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env variables")
	}
	fmt.Println("Env variables run smoothly")

	serverPort := os.Getenv("SERVER_PORT")

	dbURL := "postgres://postgres:f41zd4n11@localhost:5432/miniproject"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connection")
	}

	return db, serverPort
}

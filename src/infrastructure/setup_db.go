package infrastructure

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbName     string
	dbPort     string
)

func ConnectDB() *gorm.DB {
	loadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("connection is successful")
	}

	return db
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("fail to load .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")
}

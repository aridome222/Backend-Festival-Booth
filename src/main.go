package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbName     string
	dbPort     string
)

func main() {
	loadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("connection is successful")
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate")
	}

	user := User{
		Name:  "user1",
		Email: "user1@user.com",
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db.Create(&user)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
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

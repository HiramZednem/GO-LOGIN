package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func DBConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DSN")

	// If you want to use MYSQL use this block instead
	// and add this driver "gorm.io/driver/mysql"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}

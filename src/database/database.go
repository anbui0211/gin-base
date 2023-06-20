package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDBEcommerce() {

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	port := 5432
	sslMode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, username, password, database, port, sslMode)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("CONNECT POSTGRESQL DATABASE FAIL ... ❗️❗")
		panic(err)
	}
	log.Println("CONNECT POSTGRESQL DATABASE SUCCESSFUL ... 🚀")

	db = conn
}

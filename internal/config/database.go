package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDBEcommerce() {
	const (
		//HOST     = "localhost"
		USER     = "anbui"
		PASSWORD = "1234"
		DB_NAME  = "pg-test"
		PORT     = 5432
	)
	HOST := os.Getenv("POSTGRES_HOST")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PASSWORD, DB_NAME, PORT)
	fmt.Println("dns: ", dsn)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("CONNECT POSTGRESQL DATABASE FAIL ... ❗️❗")
		panic(err)
	}
	log.Println("CONNECT POSTGRESQL DATABASE SUCCESSFUL ... 🚀")
	db = conn
}

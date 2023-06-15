package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDBEcommerce() {
	const (
		HOST     = "localhost"
		USER     = "anbui"
		PASSWORD = "1234"
		DB_NAME  = "pg-test"
		PORT     = "5432"
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PASSWORD, DB_NAME, PORT)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("CONNECT POSTGRESQL DATABASE FAIL ... ❗️❗")
		panic(err)
	}
	log.Println("CONNECT POSTGRESQL DATABASE SUCCESSFUL ... 🚀")
	db = conn
}

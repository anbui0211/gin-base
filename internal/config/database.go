package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDBEcommerce() *gorm.DB {
	const (
		HOST     = "localhost"
		USER     = "anbui"
		PASSWORD = "1234"
		DB_NAME  = "ecommerce"
		PORT     = "5432"
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PASSWORD, DB_NAME, PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("CONNECT POSTGRESQL DATABASE FAIL ... ❗️❗")
		panic(err)
	}
	log.Println("CONNECT POSTGRESQL DATABASE SUCCESSFUL ... 🚀")
	return db
}

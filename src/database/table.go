package database

import "gorm.io/gorm"

// UserCol ...
func UserCol() *gorm.DB {
	return db.Table(userTable)
}

// ProductCol ...
func ProductCol() *gorm.DB {
	return db.Table(productTable)
}

const (
	userTable    = "users"
	productTable = "products"
)

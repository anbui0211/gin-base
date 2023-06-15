package config

import "gorm.io/gorm"

// UserCol ...
func UserCol() *gorm.DB {
	return db.Table(UserTable)
}

const (
	UserTable = "users"
)

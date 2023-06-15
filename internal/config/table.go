package config

import "gorm.io/gorm"

// UserCol ...
func UserCol() *gorm.DB {
	return ConnectDBEcommerce().Table(UserTable)
}

const (
	UserTable = "users"
)

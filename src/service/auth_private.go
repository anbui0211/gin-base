package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
	"log"
)

func (s authImpl) isExistedUser(ctx context.Context, username string) bool {
	var (
		db   = database.UserCol()
		user pgmodel.User
	)

	// Check invalid user
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false
		}

		log.Println("[Auth - isExistedUser] - err: ", err)
		return false
	}

	return true
}

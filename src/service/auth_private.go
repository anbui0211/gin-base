package service

import (
	"context"
	"gin-base/internal/config"
	pgmodel "gin-base/internal/models/pg"
	"log"
)

func (s authImpl) isExistedUser(ctx context.Context, username string) bool {
	var (
		db   = config.UserCol()
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

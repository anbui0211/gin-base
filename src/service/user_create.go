package service

import (
	"context"
	"gin-base/src/database"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
)

// Create user with client data
func (s userImpl) Create(ctx context.Context, user requestmodel.UserCreate) (res responsemodel.Upsert, err error) {
	var (
		db        = database.UserCol()
		userModel = user.ConvertToUserModel()
	)

	if err = db.Create(&userModel).Error; err != nil {
		return
	}

	res.ID = userModel.ID
	return
}

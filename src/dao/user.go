package dao

import (
	"context"
	"fmt"
	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
	querymodel "gin-base/src/model/query"
)

type UserInterface interface {
	All(c context.Context, q querymodel.UserAll) (users []pgmodel.User)
	Count(ctx context.Context) (count int64)
	Detail(ctx context.Context, id string) (user pgmodel.User, err error)
	Update(ctx context.Context, id string, payload pgmodel.User) (err error)
	ChangeStatus(ctx context.Context, id string, payload interface{}) (err error)
}

type userImpl struct{}

func User() UserInterface {
	return &userImpl{}
}

// All ...
func (userImpl) All(c context.Context, q querymodel.UserAll) (users []pgmodel.User) {
	offset := int((q.Page - 1) * q.Limit)
	limit := int(q.Limit)
	if limit <= 0 {
		limit = 5
	}

	if err := database.UserCol().Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return
	}

	return
}

// Count ...
func (userImpl) Count(ctx context.Context) (count int64) {
	if err := database.UserCol().Count(&count).Error; err != nil {
		fmt.Println("[Dao-Count] error: ", err)
		return
	}

	return
}

// Detail ...
func (userImpl) Detail(ctx context.Context, id string) (user pgmodel.User, err error) {
	if err = database.UserCol().Where("id = ?", id).First(&user).Error; err != nil {
		return
	}

	return
}

// Update ...
func (userImpl) Update(ctx context.Context, id string, payload pgmodel.User) (err error) {
	if err = database.UserCol().Where("id = ?", id).Updates(&payload).Error; err != nil {
		return
	}
	return
}

// ChangeStatus ...
func (userImpl) ChangeStatus(ctx context.Context, id string, payload interface{}) (err error) {
	if err = database.UserCol().Where("id = ?", id).Updates(&payload).Error; err != nil {
		return
	}
	return
}

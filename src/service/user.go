package service

import (
	"context"
	"gin-base/internal/config"
	pgmodel "gin-base/internal/models/pg"
	querymodel "gin-base/src/model/query"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
	"sync"
	"time"
)

type UserInterface interface {
	Create(ctx context.Context, user requestmodel.UserCreate) (res responsemodel.Upsert, err error)
	All(c context.Context, q querymodel.UserAll) (res responsemodel.UserAll)
	Detail(ctx context.Context, id string) (res *responsemodel.UserDetail, err error)
	Update(ctx context.Context, id string, payload requestmodel.UserUpdate) (res responsemodel.Upsert, err error)
	ChangeStatus(ctx context.Context, id string, payload requestmodel.UserChangeStatus) (res responsemodel.Upsert, err error)
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

func (s userImpl) Create(ctx context.Context, user requestmodel.UserCreate) (res responsemodel.Upsert, err error) {
	var (
		db        = config.UserCol()
		userModel = user.ConvertToUserModel()
	)

	if err = db.Create(&userModel).Error; err != nil {
		return
	}

	res.ID = userModel.ID.String()
	return
}

func (s userImpl) All(ctx context.Context, q querymodel.UserAll) (res responsemodel.UserAll) {

	offset := int((q.Page - 1) * q.Limit)
	limit := int(q.Limit)
	if limit <= 0 {
		limit = 5
	}

	var (
		users []pgmodel.User
		db    = config.UserCol()
	)
	if err := db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return
	}

	var wg = sync.WaitGroup{}

	wg.Add(1)
	var count int64
	go func() {
		defer wg.Done()
		db.Count(&count)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res.List = s.getListUser(ctx, users)
	}()

	wg.Wait()

	res.Total = count
	res.Limit = int64(limit)

	return
}

func (s userImpl) Detail(ctx context.Context, id string) (res *responsemodel.UserDetail, err error) {
	var (
		db   = config.UserCol()
		user pgmodel.User
	)

	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		return
	}

	res = s.detail(ctx, user)
	return
}

func (u userImpl) Update(ctx context.Context, id string, payload requestmodel.UserUpdate) (res responsemodel.Upsert, err error) {
	userUpdate := payload.ConvertToUserModel()

	var db = config.UserCol()
	if err = db.Where("id = ?", id).Updates(&userUpdate).Error; err != nil {
		return
	}

	res.ID = id
	return
}

func (u userImpl) ChangeStatus(ctx context.Context, id string, payload requestmodel.UserChangeStatus) (res responsemodel.Upsert, err error) {
	statusUpdate := map[string]interface{}{
		"status":     payload.Status,
		"created_at": time.Now(),
	}
	var db = config.UserCol()
	if err = db.Where("id = ?", id).Updates(&statusUpdate).Error; err != nil {
		return
	}

	res.ID = id
	return
}

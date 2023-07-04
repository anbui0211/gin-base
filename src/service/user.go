package service

import (
	"context"
	"errors"
	"gin-base/internal/constant"
	"gin-base/src/dao"
	"gin-base/src/database"
	querymodel "gin-base/src/model/query"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
	"sync"
	"time"
)

type UserInterface interface {
	Create(ctx context.Context, user requestmodel.UserCreate) (res responsemodel.Upsert, err error)
	All(c context.Context, q querymodel.UserAll) (res responsemodel.UserAll)
	Detail(ctx context.Context, id string) (res responsemodel.UserDetail, err error)
	Update(ctx context.Context, id string, payload requestmodel.UserUpdate) (res *responsemodel.Upsert, err error)
	ChangeStatus(ctx context.Context, id string, payload requestmodel.UserChangeStatus) (res responsemodel.Upsert, err error)
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

// Create ...
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

// All ...
func (s userImpl) All(ctx context.Context, q querymodel.UserAll) (res responsemodel.UserAll) {

	var (
		wg = sync.WaitGroup{}
		d  = dao.User()
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		users := d.All(ctx, q)
		res.List = s.getListUser(ctx, users)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res.Total = d.Count(ctx)
	}()

	wg.Wait()

	res.Limit = q.Limit

	return
}

// Detail ...
func (s userImpl) Detail(ctx context.Context, id string) (res responsemodel.UserDetail, err error) {
	var d = dao.User()
	user, err := d.FindByID(ctx, id)
	if err != nil {
		return
	}

	res = s.detail(ctx, user)
	return
}

// Update ...
func (u userImpl) Update(ctx context.Context, id string, payload requestmodel.UserUpdate) (res *responsemodel.Upsert, err error) {
	var (
		d          = dao.User()
		userUpdate = payload.ConvertToUserModel()
	)

	user, err := d.FindByID(ctx, id)
	if err != nil || err.Error() == constant.ErrRecordNotFound {

	}

	if err = d.Update(ctx, id, userUpdate); err != nil {
		return nil, errors.New("error when change user")
	}

	res.ID = user.ID
	return
}

// ChangeStatus ...
func (u userImpl) ChangeStatus(ctx context.Context, id string, payload requestmodel.UserChangeStatus) (res responsemodel.Upsert, err error) {
	var (
		d            = dao.User()
		statusUpdate = map[string]interface{}{
			"status":     payload.Status,
			"created_at": time.Now(),
		}
	)

	user, err := d.FindByID(ctx, id)
	if err != nil || err.Error() == constant.ErrRecordNotFound {

	}

	if err = d.ChangeStatus(ctx, id, statusUpdate); err != nil {
		err = errors.New("errorcode when change user")
		return
	}

	res.ID = user.ID
	return
}

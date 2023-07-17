package service

import (
	"context"
	"errors"
	"gin-base/internal/constant"
	"gin-base/src/dao"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
	"time"
)

// Update ...
func (u userImpl) Update(ctx context.Context, id string, payload requestmodel.UserUpdate) (res *responsemodel.Upsert, err error) {
	var (
		d          = dao.User()
		userUpdate = payload.ConvertToUserModel()
	)

	user, err := d.FindByID(ctx, id)
	if err != nil || err.Error() == constant.ErrRecordNotFound {
		return
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
		return
	}

	if err = d.ChangeStatus(ctx, id, statusUpdate); err != nil {
		err = errors.New("errorcode when change user")
		return
	}

	res.ID = user.ID
	return
}

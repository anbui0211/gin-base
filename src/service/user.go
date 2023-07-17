package service

import (
	"context"
	querymodel "gin-base/src/model/query"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
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

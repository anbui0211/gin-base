package service

import (
	"context"
	"fmt"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
)

type AuthInterface interface {
	Register(ctx context.Context, payload requestmodel.Register) (responsemodel.Auth, error)
}

type authImpl struct{}

func Auth() AuthInterface {
	return authImpl{}
}

func (a authImpl) Register(ctx context.Context, payload requestmodel.Register) (responsemodel.Auth, error) {
	fmt.Println("payload: ", payload)
	fmt.Println("REGISTER IN SERVICE ...")
	return responsemodel.Auth{}, nil
}

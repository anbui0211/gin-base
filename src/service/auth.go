package service

import (
	"context"
	"errors"
	authinternal "gin-base/internal/auth"
	"gin-base/internal/config"
	"gin-base/internal/constant"
	pgmodel "gin-base/internal/models/pg"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
)

type AuthInterface interface {
	Register(ctx context.Context, payload requestmodel.Register) (*responsemodel.Auth, error)
	Login(ctx context.Context, payload requestmodel.Login) (*responsemodel.Auth, error)
}

type authImpl struct{}

func Auth() AuthInterface {
	return authImpl{}
}

func (s authImpl) Register(ctx context.Context, payload requestmodel.Register) (*responsemodel.Auth, error) {
	var db = config.UserCol()

	// Check user exist
	if isExisted := s.isExistedUser(ctx, payload.Username); isExisted {
		return nil, errors.New(constant.ErrAlreadyExistUsername)
	}

	newUser := payload.ConvertToModel()

	// Hash password
	hashPass, err := authinternal.HashPassword(newUser.Password)
	if err != nil {
		return nil, errors.New(constant.ErrAuthHashPassword)
	}
	newUser.Password = hashPass

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	// generate token
	token, err := authinternal.GenerateToken(authinternal.User{
		ID:   newUser.ID.String(),
		Name: newUser.Name,
	})

	if err != nil {
		return nil, err
	}

	return &responsemodel.Auth{
		Token: token,
	}, nil
}

func (s authImpl) Login(ctx context.Context, payload requestmodel.Login) (*responsemodel.Auth, error) {

	var (
		db   = config.UserCol()
		user pgmodel.User
	)

	if err := db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return nil, errors.New(constant.ErrUsernameNotExist)
	}

	if err := authinternal.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, errors.New(constant.ErrAuthInvalidPassword)
	}

	token, err := authinternal.GenerateToken(authinternal.User{
		ID:   user.ID.String(),
		Name: user.Name,
	})

	if err != nil {
		return nil, err
	}

	return &responsemodel.Auth{Token: token}, nil
}

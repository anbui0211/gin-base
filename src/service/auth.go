package service

import (
	"context"
	"errors"
	authinternal "gin-base/internal/auth"
	"gin-base/src/errorcode"

	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
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
	var db = database.UserCol()

	// Check user exist
	if isExisted := s.isExistedUser(ctx, payload.Username); isExisted {
		return nil, errors.New(errorcode.ErrAlreadyExistUsername)
	}

	newUser := payload.ConvertToModel()

	// Hash password
	hashPass, err := authinternal.HashPassword(newUser.Password)
	if err != nil {
		return nil, errors.New(errorcode.ErrAuthHashPassword)
	}
	newUser.Password = hashPass

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	// generate token
	token, err := authinternal.GenerateToken(authinternal.User{
		ID:   newUser.ID,
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
		db   = database.UserCol()
		user pgmodel.User
	)

	if err := db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return nil, errors.New(errorcode.ErrUsernameNotExist)
	}

	if err := authinternal.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, errors.New(errorcode.ErrAuthInvalidPassword)
	}

	token, err := authinternal.GenerateToken(authinternal.User{
		ID:   user.ID,
		Name: user.Name,
	})

	if err != nil {
		return nil, err
	}

	return &responsemodel.Auth{Token: token}, nil
}

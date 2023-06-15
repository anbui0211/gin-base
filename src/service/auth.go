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
	"github.com/google/uuid"
	"time"
)

type AuthInterface interface {
	Register(ctx context.Context, payload requestmodel.Register) (*responsemodel.Auth, error)
}

type authImpl struct{}

func Auth() AuthInterface {
	return authImpl{}
}

func (s authImpl) Register(ctx context.Context, payload requestmodel.Register) (*responsemodel.Auth, error) {
	var (
		db   = config.UserCol()
		user pgmodel.User
	)

	// Check invalid user
	if err := db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New(constant.ErrIsExistedUsername)
		}

		return nil, err
	}

	newUser := pgmodel.User{
		PgModel: pgmodel.PgModel{
			ID:        uuid.New(),
			Status:    constant.StatusInactive,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: payload.Username,
		Password: payload.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	token, err := authinternal.GenerateToken(authinternal.User{
		ID:   user.ID.String(),
		Name: user.Name,
	})

	if err != nil {
		return nil, err
	}

	return &responsemodel.Auth{
		Token: token,
	}, nil
}

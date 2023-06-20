package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	responsemodel "gin-base/src/model/response"
	"sync"
)

// getListUser ...
func (s userImpl) getListUser(ctx context.Context, users []pgmodel.User) []responsemodel.UserBrief {
	var (
		wg     sync.WaitGroup
		result = make([]responsemodel.UserBrief, len(users))
	)

	wg.Add(len(users))
	for index, user := range users {
		go func(i int, u pgmodel.User) {
			defer wg.Done()
			result[i] = s.brief(ctx, u)
		}(index, user)
	}

	wg.Wait()
	return result
}

// brief
func (s userImpl) brief(ctx context.Context, user pgmodel.User) responsemodel.UserBrief {
	return responsemodel.UserBrief{
		ID:        user.ID.String(),
		Name:      user.Name,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// detail ...
func (s userImpl) detail(ctx context.Context, user pgmodel.User) *responsemodel.UserDetail {
	return &responsemodel.UserDetail{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/dao"
)

// FindByID ...
func (productImpl) FindByID(ctx context.Context, id string) (product pgmodel.Product, err error) {
	return dao.Product().FindByID(ctx, id)
}

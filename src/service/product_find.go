package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/dao"
	querymodel "gin-base/src/model/query"
)

func (s *productImpl) FindByID(ctx context.Context, id string) (product pgmodel.Product, err error) {
	return dao.Product().FindByID(ctx, id)
}

func (s *productImpl) FindByCondition(ctx context.Context, q querymodel.Product) []pgmodel.Product {
	return dao.Product().FindByCondition(ctx, q)
}

func (s *productImpl) FindRowsByCondition(ctx context.Context, q querymodel.Product) (result []pgmodel.Product, err error) {
	return dao.Product().FindRowsByCondition(ctx, q)
}

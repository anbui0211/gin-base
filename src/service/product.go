package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/dao"
)

type ProductInterface interface {
	FindByID(ctx context.Context, id string) (product pgmodel.Product, err error)
	CreateManyProduct(ctx context.Context, products []pgmodel.Product) error
}

type productImpl struct {
}

func Product() ProductInterface {
	return &productImpl{}
}

// FindByID ...
func (productImpl) FindByID(ctx context.Context, id string) (product pgmodel.Product, err error) {
	return dao.Product().FindByID(ctx, id)
}

// CreateManyProduct ...
func (productImpl) CreateManyProduct(ctx context.Context, products []pgmodel.Product) error {
	return dao.Product().InsertMany(ctx, products)
}

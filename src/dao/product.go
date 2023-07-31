package dao

import (
	"context"

	pgmodel "gin-base/internal/models"
	querymodel "gin-base/src/model/query"
)

type ProductInterface interface {
	FindByID(ctx context.Context, id string) (product pgmodel.Product, err error)
	BulkCreate(ctx context.Context, products []pgmodel.Product) error
	FindByCondition(ctx context.Context, q querymodel.Product) (result []pgmodel.Product)
	FindRowsByCondition(ctx context.Context, q querymodel.Product) (result []pgmodel.Product, err error)
}

type productImpl struct {
}

func Product() ProductInterface {
	return &productImpl{}
}

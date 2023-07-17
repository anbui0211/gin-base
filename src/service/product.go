package service

import (
	"context"
	pgmodel "gin-base/internal/models"
)

type ProductInterface interface {
	FindByID(ctx context.Context, id string) (product pgmodel.Product, err error)
	BulkCreateProduct(ctx context.Context, products []pgmodel.Product) error
}

type productImpl struct {
}

func Product() ProductInterface {
	return &productImpl{}
}

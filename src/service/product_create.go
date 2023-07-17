package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/dao"
)

// BulkCreateProduct ...
func (productImpl) BulkCreateProduct(ctx context.Context, products []pgmodel.Product) error {
	return dao.Product().InsertMany(ctx, products)
}

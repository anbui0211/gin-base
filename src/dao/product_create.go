package dao

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
	"log"
)

func (productImpl) BulkCreate(ctx context.Context, products []pgmodel.Product) error {
	if err := database.ProductCol().Create(products).Error; err != nil {
		log.Fatal("[product-dao-insertMany] error: ", err)
		return err
	}

	return nil
}

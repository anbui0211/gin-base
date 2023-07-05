package dao

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
	"log"
)

type ProductInterface interface {
	FindByID(ctx context.Context, id string) (product pgmodel.Product, err error)
	InsertMany(ctx context.Context, products []pgmodel.Product) error
}

type productImpl struct {
}

func Product() ProductInterface {
	return &productImpl{}
}

// FindByID ...
func (productImpl) FindByID(ctx context.Context, id string) (product pgmodel.Product, err error) {
	if err = database.ProductCol().Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("[Product-dao-findByID] error: ", err)
		return
	}

	return
}

func (productImpl) InsertMany(ctx context.Context, products []pgmodel.Product) error {
	if err := database.ProductCol().Create(products).Error; err != nil {
		log.Fatal("[product-dao-insertMany] error: ", err)
		return err
	}

	return nil
}

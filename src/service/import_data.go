package service

import (
	"context"
	pgmodel "gin-base/internal/models"
)

type ImportDataInterface interface {
	ImportProducts(ctx context.Context, fileName string) (int, error)
}

type importDataImpl struct {
	rs       ImportRowSet
	products []pgmodel.Product
}

func ImportData() ImportDataInterface {
	return &importDataImpl{}
}

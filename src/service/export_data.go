package service

import (
	"context"
	querymodel "gin-base/src/model/query"
)

type ExportDataInterface interface {
	ExportProduct(ctx context.Context, q querymodel.Product) error
}

type exportDataImpl struct{}

func ExportData() ExportDataInterface {
	return &exportDataImpl{}
}

// ExportProduct ...
func (s exportDataImpl) ExportProduct(ctx context.Context, q querymodel.Product) error {
	return nil
}

package service

import (
	"context"
	"fmt"
	"gin-base/src/errorcode"
	querymodel "gin-base/src/model/query"

	"gopkg.in/errgo.v2/fmt/errors"
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
	fmt.Println("query: ", q)

	var (
		pSvc = Product()
	)

	products, err := pSvc.FindRowsByCondition(ctx, q)
	if err != nil {
		return errors.New(errorcode.ErrFindDataProduct)
	}

	fmt.Println(products)
	// TODO: Tiếp tục làm tại đây (viết ra file PDF)

	return nil
}

package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	"gin-base/src/errorcode"
	"gopkg.in/errgo.v2/errors"
	"io"
	"log"
	"os"
)

const (
	ImportColumnProductId           = 0
	ImportColumnProductName         = 1
	ImportColumnProductSearchString = 2
	ImportColumnProductCategoryId   = 3
	ImportColumnProductQuantity     = 4
	ImportColumnProductPrice        = 5
	ImportColumnProductStatus       = 6
)

var ProductImportDef = map[int]ImportColumDef{
	ImportColumnProductId:           {Name: "id", Def: regexStringDigit},
	ImportColumnProductName:         {Name: "name", Def: regexRequireAlphaVn},
	ImportColumnProductSearchString: {Name: "searchString", Def: regexRequireAlpha},
	ImportColumnProductCategoryId:   {Name: "categoryId", Def: regexRequireDigit},
	ImportColumnProductQuantity:     {Name: "quantity", Def: regexRequireDigit},
	ImportColumnProductPrice:        {Name: "price", Def: regexDecimal},
	ImportColumnProductStatus:       {Name: "status", Def: regexRequireAlphaVn},
}

// ImportProducts ...
func (s *importDataImpl) ImportProducts(ctx context.Context, fileName string) (int, error) {
	var (
		appFolder = "/assets/data_csv"
	)

	curDir, _ := os.Getwd()               // ../gin-base
	curDir += appFolder                   // ../gin-base/assets/data_csv
	importFile := curDir + "/" + fileName // ../gin-base/assets/data_csv/products.csv

	f, err := os.Open(importFile)
	if err != nil {
		log.Println("Error open file: ", err)
		return 0, errors.New(errorcode.ErrOpenFile)

	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("Error close file: ", err)
		}
	}()

	if err := s.importDataProduct(ctx, f); err != nil {
		return 0, err
	}

	return len(s.products), nil
}

// importDataProduct ...
func (s *importDataImpl) importDataProduct(ctx context.Context, data io.ReadCloser) (err error) {
	// ReadAndCheckCsv
	if err = s.readAndCheckCsv(ctx, data, ProductImportDef, &s.rs); err != nil {
		return
	}

	// Set items
	if err := s.setItemsProduct(ctx); err != nil {
	}

	var pSvc = Product()
	if err = pSvc.CreateManyProduct(ctx, s.products); err != nil {
		return
	}

	return
}

// setItemsProduct ...
func (s *importDataImpl) setItemsProduct(ctx context.Context) error {
	for _, row := range s.rs.rows {
		product, err := s.setProduct(ctx, row)
		if err != nil {
			return err
		}
		s.products = append(s.products, product)
	}

	return nil
}

// getProductId ...
func getProductId(row *ImportRow) (productId string) {
	productId = row.String(ImportColumnProductId)
	return
}

// getProductName ...
func getProductName(row *ImportRow) (productName string) {
	productName = row.String(ImportColumnProductName)
	return
}

// setProduct ...
func (s *importDataImpl) setProduct(ctx context.Context, row *ImportRow) (product pgmodel.Product, err error) {
	// handle empty and invalid

	// Set product
	product.ID = row.String(ImportColumnProductId)
	product.Name = row.String(ImportColumnProductName)
	product.SearchString = row.String(ImportColumnProductSearchString)
	product.CategoryID = row.String(ImportColumnProductCategoryId)
	product.Quantity = row.Int64(ImportColumnProductQuantity)
	product.Price = row.Float64(ImportColumnProductPrice)
	product.Status = row.String(ImportColumnProductStatus)

	return
}

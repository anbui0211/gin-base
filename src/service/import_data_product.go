package service

import (
	"context"
	"io"
	"log"
	"os"

	pgmodel "gin-base/internal/models"
	"gin-base/src/errorcode"

	"gopkg.in/errgo.v2/errors"
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
		appFolder = "/assets/testdata"
	)

	curDir, _ := os.Getwd()               // ../gin-base
	curDir += appFolder                   // ../gin-base/assets/testdata
	importFile := curDir + "/" + fileName // ../gin-base/assets/testdata/products.csv

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

func (s *importDataImpl) importDataProduct(ctx context.Context, data io.ReadCloser) (err error) {
	// ReadAndCheckCsv
	if err = s.readAndCheckCsv(ctx, data, ProductImportDef, &s.rs); err != nil {
		return
	}

	// Set items
	if err := s.setItemsProduct(ctx); err != nil {
		return err
	}

	var pSvc = Product()
	if err = pSvc.BulkCreateProduct(ctx, s.products); err != nil {
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

// setProduct ...
func (s *importDataImpl) setProduct(ctx context.Context, row *ImportRow) (product pgmodel.Product, err error) {
	// handle empty and invalid

	// Set product
	product.ID = row.toString(ImportColumnProductId)
	product.Name = row.toString(ImportColumnProductName)
	product.SearchString = row.toString(ImportColumnProductSearchString)
	product.CategoryID = row.toString(ImportColumnProductCategoryId)
	product.Quantity = row.toInt64(ImportColumnProductQuantity)
	product.Price = row.toFloat64(ImportColumnProductPrice)
	product.Status = row.toString(ImportColumnProductStatus)

	return
}

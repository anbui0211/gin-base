package service

import (
	"context"
	"gin-base/src/errorcode"
	"gopkg.in/errgo.v2/errors"
	"io"
	"log"
	"os"
)

//
// Product ...
//

const (
	ImportColumnProductName         = 0
	ImportColumnProductSearchString = 1
	ImportColumnProductCategoryId   = 2
	ImportColumnProductQuantity     = 3
	ImportColumnProductPrice        = 4
	ImportColumnProductStatus       = 5
)

var ProductImportDef = map[int]ImportColumDef{
	ImportColumnProductName:         {Name: "name", Def: regexRequireAlphaVn},
	ImportColumnProductSearchString: {Name: "searchString", Def: regexRequireAlpha},
	ImportColumnProductCategoryId:   {Name: "categoryId", Def: regexRequireDigit},
	ImportColumnProductQuantity:     {Name: "quantity", Def: regexRequireDigit},
	ImportColumnProductPrice:        {Name: "price", Def: regexDecimal},
	ImportColumnProductStatus:       {Name: "status", Def: regexRequireAlphaVn},
}

func (s importDataImpl) ImportProducts(ctx context.Context, fileName string) error {
	var (
		appFolder = "/assets/data_csv"
	)

	curDir, _ := os.Getwd()               // ../gin-base
	curDir += appFolder                   // ../gin-base/assets/data_csv
	importFile := curDir + "/" + fileName // ../gin-base/assets/data_csv/products.csv

	f, err := os.Open(importFile)
	if err != nil {
		log.Println("Error open file: ", err)
		return errors.New(errorcode.ErrOpenFile)

	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Error open file: ", err)
		}
	}()

	if err := s.importDataProduct(ctx, f); err != nil {
		return err
	}

	return nil
}

func (s importDataImpl) importDataProduct(ctx context.Context, data io.ReadCloser) (err error) {
	// ReadAndCheckCsv
	if err = s.readAndCheckCsv(ctx, data, ProductImportDef, &s.rs); err != nil {
		return
	}

	// Set items
	if err := s.setItemsProduct(ctx); err != nil {
	}
	return
}

func (s importDataImpl) setItemsProduct(ctx context.Context) (err error) {

	return
}

// getProductName ...
func getProductName(row *ImportRow) (productName string) {
	productName = row.String(ImportColumnProductName)
	return
}

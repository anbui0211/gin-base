package service

import (
	"context"
)

type ImportDataInterface interface {
	ImportProducts(ctx context.Context, fileName string) error
}

type importDataImpl struct {
	rs ImportRowSet
	//items []AccountBizItem
}

func ImportData() ImportDataInterface {
	return &importDataImpl{}
}

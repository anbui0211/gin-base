package service

type ImportDataItemInterface interface {
}

type importDataItemImpl struct {
	rs ImportRowSet
	//items []AccountBizItem
}

func ImportDataItem() ImportDataItemInterface {
	return &importDataItemImpl{}
}

type importDataProductItemForListImpl struct {
}

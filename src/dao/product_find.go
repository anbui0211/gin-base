package dao

import (
	"context"
	"log"

	pgmodel "gin-base/internal/models"
	"gin-base/src/database"
	querymodel "gin-base/src/model/query"
)

// FindByID ...
func (productImpl) FindByID(ctx context.Context, id string) (product pgmodel.Product, err error) {
	if err = database.ProductCol().Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("[Product-dao-findByID] error: ", err)
		return
	}

	return
}

// FindByCondition ...
func (productImpl) FindByCondition(ctx context.Context, q querymodel.Product) (result []pgmodel.Product) {
	offset := int((q.Page - 1) * q.Limit)
	limit := int(q.Limit)
	if limit <= 0 {
		limit = 5
	}

	pQuery := &pgmodel.Product{}
	var pArgs []interface{}
	if q.CategoryID != "" {
		pQuery.CategoryID = q.CategoryID
		pArgs = append(pArgs, "CategoryID")
	}
	if q.Status != "" {
		pQuery.Status = q.Status
		pArgs = append(pArgs, "CategoryID")
	}

	if err := database.ProductCol().
		Where(pQuery, pArgs...).
		Limit(limit).
		Offset(offset).
		Find(&result).Error; err != nil {
		return
	}

	return
}

// FindRowsByCondition ...
func (productImpl) FindRowsByCondition(ctx context.Context, q querymodel.Product) (result []pgmodel.Product, err error) {
	offset := int((q.Page - 1) * q.Limit)
	limit := int(q.Limit)
	if limit <= 0 {
		limit = 5
	}

	pQuery := &pgmodel.Product{}
	var pArgs []interface{}
	if q.CategoryID != "" {
		pQuery.CategoryID = q.CategoryID
		pArgs = append(pArgs, "CategoryID")
	}
	if q.Status != "" {
		pQuery.Status = q.Status
		pArgs = append(pArgs, "CategoryID")
	}

	db := database.ProductCol()

	rows, err := db.
		Where(pQuery, pArgs...).
		Limit(limit).
		Offset(offset).
		Rows()

	if err != nil {
		return result, err
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		count++

		var p pgmodel.Product
		err = db.ScanRows(rows, &p)
		if err != nil {
			return
		}

		result = append(result, p)
	}

	return
}

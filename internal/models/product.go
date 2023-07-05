package pgmodel

type Product struct {
	PgModel      `gorm:",inline"`
	Name         string  `gorm:"column:name"`
	SearchString string  `gorm:"column:search_string"`
	CategoryID   string  `gorm:"column:category_id"`
	Quantity     int64   `gorm:"column:quantity"`
	Price        float64 `gorm:"column:price"`
	Status       string  `gorm:"column:status"`
}

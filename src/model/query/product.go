package querymodel

import validation "github.com/go-ozzo/ozzo-validation"

type Product struct {
	Page       int64  `form:"page"`
	Limit      int64  `form:"limit"`
	CategoryID string `form:"category_id"`
	Status     string `form:"status"`
}

func (m Product) Validate() error {
	return validation.ValidateStruct(&m)
}

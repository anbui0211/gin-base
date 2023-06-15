package querymodel

import validation "github.com/go-ozzo/ozzo-validation"

type UserAll struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func (m UserAll) Validate() error {
	return validation.ValidateStruct(&m)
}

package requestmodel

import (
	"gin-base/internal/constant"
	pgmodel "gin-base/internal/models/pg"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (m Register) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required.Error(constant.ErrEmptyUsername)),
		validation.Field(&m.Password, validation.Required.Error(constant.ErrEmptyPassword)),
		validation.Field(&m.Phone, validation.Required.Error(constant.ErrEmptyPassword)),
		validation.Field(&m.Email, validation.Required.Error(constant.ErrEmptyPassword)),
	)
}

func (m Register) ConvertToModel() pgmodel.User {
	return pgmodel.User{
		Username: m.Username,
		Password: m.Password,
	}
}

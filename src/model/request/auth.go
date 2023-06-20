package requestmodel

import (
	"gin-base/internal/constant"
	pgmodel "gin-base/internal/models"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (m Register) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required.Error(constant.ErrEmptyUsername)),
		validation.Field(&m.Password, validation.Required.Error(constant.ErrEmptyPassword)),
		validation.Field(&m.Phone, validation.Required.Error(constant.ErrEmptyPhone)),
		validation.Field(&m.Email, validation.Required.Error(constant.ErrEmptyEmail)),
		validation.Field(&m.Name, validation.Required.Error(constant.ErrNameNotExist)),
	)
}

func (m Register) ConvertToModel() pgmodel.User {
	return pgmodel.User{
		PgModel: pgmodel.PgModel{
			ID:        uuid.New(),
			Status:    "inactive",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: m.Username,
		Password: m.Password,
		Phone:    m.Phone,
		Email:    m.Email,
		Name:     m.Name,
	}

}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m Login) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required.Error(constant.ErrEmptyUsername)),
		validation.Field(&m.Password, validation.Required.Error(constant.ErrEmptyPassword)),
	)
}

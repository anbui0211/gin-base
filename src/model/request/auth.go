package requestmodel

import (
	pgmodel "gin-base/internal/models"
	"gin-base/src/errorcode"
	"github.com/google/uuid"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
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
		validation.Field(&m.Username, validation.Required.Error(errorcode.ErrEmptyUsername)),
		validation.Field(&m.Password, validation.Required.Error(errorcode.ErrEmptyPassword)),
		validation.Field(&m.Phone, validation.Required.Error(errorcode.ErrEmptyPhone)),
		validation.Field(&m.Email, validation.Required.Error(errorcode.ErrEmptyEmail)),
		validation.Field(&m.Name, validation.Required.Error(errorcode.ErrNameNotExist)),
	)
}

func (m Register) ConvertToModel() pgmodel.User {
	return pgmodel.User{
		PgModel: pgmodel.PgModel{
			//ID:        uuid.New(),
			Status:    "inactive",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: m.Username,
		Password: m.Password,
		Name:     m.Name,
		Phone:    m.Phone,
		Email:    m.Email,
		UserID:   uuid.New().String(),
	}

}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m Login) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required.Error(errorcode.ErrEmptyUsername)),
		validation.Field(&m.Password, validation.Required.Error(errorcode.ErrEmptyPassword)),
	)
}

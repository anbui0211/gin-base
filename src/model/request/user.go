package requestmodel

import (
	"gin-base/internal/constant"
	pgmodel "gin-base/internal/models/pg"
	updatemodel "gin-base/src/model/update"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"time"
)

type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (m UserCreate) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required.Error("email cannot be empty")),
		validation.Field(&m.Name, validation.Required.Error("name cannot be empty")),
	)
}

func (m UserCreate) ConvertToUserModel() pgmodel.User {
	return pgmodel.User{
		PgModel: pgmodel.PgModel{
			ID:        uuid.New(),
			Status:    "inactive",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  m.Name,
		Email: m.Email,
	}
}

type UserUpdate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (m UserUpdate) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required.Error("email cannot be empty")),
		validation.Field(&m.Name, validation.Required.Error("name cannot be empty")),
	)
}

func (m UserUpdate) ConvertToUserModel() updatemodel.UserUpdate {
	return updatemodel.UserUpdate{
		Name:      m.Name,
		Email:     m.Email,
		UpdatedAt: time.Now(),
	}
}

type UserChangeStatus struct {
	Status string `json:"status"`
}

func (m UserChangeStatus) Validate() error {
	var status = []interface{}{
		constant.StatusActive,
		constant.StatusInactive,
	}

	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.In(status...).Error(constant.ErrStatusInvalid)),
	)
}

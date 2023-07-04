package pgmodel

import (
	"time"
)

type PgModel struct {
	ID        string    `gorm:"column:id"`
	Status    string    `gorm:"column:status;not null;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

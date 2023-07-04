package pgmodel

import (
	"time"
)

type PgModel struct {
	ID        int       `gorm:"column:id;auto_increment;primary_key;not null"`
	Status    string    `gorm:"column:status;not null;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

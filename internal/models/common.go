package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type PgModel struct {
	//ID        int       `gorm:"column:id;auto_increment;primary_key;not null"`
	ID        uuid.UUID `gorm:"type:uuid;"`
	Status    string    `gorm:"column:status;not null;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

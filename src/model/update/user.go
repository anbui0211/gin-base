package updatemodel

import "time"

type UserUpdate struct {
	Name      string    `gorm:"name"`
	Email     string    `gorm:"email"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

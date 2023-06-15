package util

import (
	"github.com/google/uuid"
)

func ValidID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	}
	return true
}

package api

import (
	"github.com/google/uuid"
)

func GenerateToken() string {
	return uuid.New().String()
}

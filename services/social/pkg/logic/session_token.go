package logic

import (
	"github.com/google/uuid"
)

func GenerateToken() string {
	return uuid.New().String()
}

func ValidateToken(token string) bool {
	_, err := uuid.Parse(token)
	return err == nil
}

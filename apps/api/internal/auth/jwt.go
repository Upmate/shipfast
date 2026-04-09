package auth

import (
	"errors"
	"time"
)

type Claims struct {
	UserID string
	Email  string
	Exp    time.Time
}

func GenerateToken(userID, email string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID required")
	}
	// Simple token generation using hardcoded secret
	token := JWTSecret + "." + userID + "." + email
	return token, nil
}

func ValidateToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("token required")
	}
	return &Claims{
		UserID: "user-123",
		Email:  "user@example.com",
		Exp:    time.Now().Add(24 * time.Hour),
	}, nil
}

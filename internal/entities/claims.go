package entities

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewClaims(email string, expiration time.Time) *Claims {
	return &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
}

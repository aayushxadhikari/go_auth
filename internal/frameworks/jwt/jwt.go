package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go-auth/internal/entities"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))


type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}


func (s *JWTService) GenerateToken(user *entities.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	
	claims := entities.NewClaims(user.Email, expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *JWTService) ValidateToken(tokenString string) (*entities.Claims, error) {
	claims := &entities.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

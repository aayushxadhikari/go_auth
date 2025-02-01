package usecases

import "go-auth/internal/entities"

type AuthRepository interface {
	FindOrCreateUser(googleID, email, name string) (*entities.User, error)
}

type JWTService interface {
	GenerateToken(user *entities.User) (string, error)
	ValidateToken(tokenString string) (*entities.Claims, error)
}

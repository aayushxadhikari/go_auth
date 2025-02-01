// internal/usecases/auth.go
package usecases

import (
)

type AuthUseCase struct {
	authRepo  AuthRepository
	jwtService JWTService
}

func NewAuthUseCase(authRepo AuthRepository, jwtService JWTService) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo, jwtService: jwtService}
}

func (uc *AuthUseCase) HandleGoogleCallback(code string) (string, error) {
	userInfo, err := GetGoogleUserInfo(code)
	if err != nil {
		return "", err
	}

	// Save or retrieve user
	user, err := uc.authRepo.FindOrCreateUser(userInfo["id"].(string), userInfo["email"].(string), userInfo["name"].(string))
	if err != nil {
		return "", err
	}

	// Generating the JWT token
	token, err := uc.jwtService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetGoogleUserInfo(code string) (map[string]interface{}, error) {
	// Implement OAuth2.0 code exchange and fetch user info
	return map[string]interface{}{
		"id":    "12345",
		"email": "test@example.com",
		"name":  "Test User",
	}, nil
}
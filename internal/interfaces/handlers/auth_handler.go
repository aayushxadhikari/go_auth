// internal/interfaces/handlers/auth_handler.go
package handlers

import (
	"net/http"
	"go-auth/internal/usecases"
)

type AuthHandler struct {
	authUseCase *usecases.AuthUseCase
}

func NewAuthHandler(authUseCase *usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect to Google OAuth login
}

func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := h.authUseCase.HandleGoogleCallback(code)
	if err != nil {
		http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func (h *AuthHandler) Protected(w http.ResponseWriter, r *http.Request) {
	// Validate JWT token and return protected data
}
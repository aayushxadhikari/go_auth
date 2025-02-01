// internal/entities/user.go
package entities

type User struct {
	ID       uint   `json:"id"`
	GoogleID string `json:"google_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
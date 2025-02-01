package repositories

import (
	"go-auth/internal/entities"
	"gorm.io/gorm" 
)

type UserRepository struct {
	db *gorm.DB 
}

func NewUserRepository(db *gorm.DB) *UserRepository { 
	return &UserRepository{db: db}
}

func (r *UserRepository) FindOrCreateUser(googleID, email, name string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("google_id = ?", googleID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user = entities.User{GoogleID: googleID, Email: email, Name: name}
		if err := r.db.Create(&user).Error; err != nil {
			return nil, err
		}
	}
	return &user, err
}

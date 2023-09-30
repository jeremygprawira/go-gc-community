package repository

import (
	"go-gc-community/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user *model.User
    err := r.db.Where("email = ?", email).Find(&user).Error
    if err != nil {
        return user, err
    }

    return user, nil
}
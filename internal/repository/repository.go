package repository

import (
	"go-gc-community/internal/model"

	"gorm.io/gorm"
)

type Users interface {
	Create(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type Repositories struct {
	Users	Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
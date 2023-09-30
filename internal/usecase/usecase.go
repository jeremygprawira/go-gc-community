package usecase

import (
	"go-gc-community/internal/model"
	"go-gc-community/internal/repository"
)

type RegisterUserRequest struct {
	Name		string		`json:"name" binding:"required"`
	Email	 	string		`json:"email" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}

type Users interface {
	Register(request RegisterUserRequest) (*model.User, error)
}

type Usecases struct {
	Users	Users
}

type Dependencies struct {
	Repository	*repository.Repositories
}

func NewUsecases(dependency Dependencies) *Usecases {
	userUsecase := NewUserUsecase(dependency.Repository.Users)

	return &Usecases{
		Users: userUsecase,
	}
}
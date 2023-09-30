package usecase

import (
	"errors"
	"go-gc-community/internal/model"
	"go-gc-community/internal/repository"
	"go-gc-community/pkg/email"
	"go-gc-community/pkg/hash"
)

type UserService struct {
	repository	repository.Users
}

func NewUserUsecase(repo repository.Users) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) Register(request RegisterUserRequest) (*model.User, error){
	user := model.User{}
	user.Name = request.Name
	
	inputEmail := request.Email
	isEmailValid, err := email.IsEmailValid(inputEmail)
	if !isEmailValid {
		return &user, err
	}

	isEmailAvailable, err := s.repository.FindByEmail(inputEmail)
	if err != nil {
		return &user, err
	}

	if isEmailAvailable.ID != 0 {
		return &user, errors.New("user with this e-mail is already available")
	}

	user.Email = inputEmail

	encrypted, err := hash.Encyrpt(request.Password)
	if err != nil {
		return &user, err
	}

	user.Password = encrypted
	user.Role = "Jemaat"
	user.IsCoolMember = false

	newUser, err := s.repository.Create(&user)
	if err != nil {
		return &user, err
	}

	return newUser, nil
}
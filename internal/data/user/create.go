package user

import (
	"errors"
	"fmt"

	"github.com/tomazcx/go-chat-app/internal/entity"
	"github.com/tomazcx/go-chat-app/internal/infra/database/userdb"
)

var (
	ErrLoginAlreadyRegistered = errors.New("Login already registered")
	ErrPasswordsDoNotMatch = errors.New("Passwords do not match")
)

type CreateUserDTO struct {
	Login           string
	Password        string
	ConfirmPassword string
}

type CreateUserUseCase struct {
	repo *userdb.UserRepository
}

func (uc *CreateUserUseCase) Execute(dto CreateUserDTO) (*entity.User, error) {
	
	_, err := uc.repo.FindByLogin(dto.Login)

	if err == nil {
		return nil, ErrLoginAlreadyRegistered
	}

	if dto.Password != dto.ConfirmPassword {
		return nil, ErrPasswordsDoNotMatch
	}

	user, err := entity.NewUser(dto.Login, dto.Password)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating the user: %v", err))
	}

	err = uc.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil	
}

func NewCreateUserUseCase() *CreateUserUseCase {
	repo := userdb.NewUserRepository()
	return &CreateUserUseCase{
		repo: repo,
	}
}

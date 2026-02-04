package service

import (
	"context"
	"errors"
	"strings"

	"github.com/cdlinkin/system-booking/internal/model"
	"github.com/cdlinkin/system-booking/internal/repo"
)

type UserService interface {
	Register(ctx context.Context, input *RegisterUserDTO) (*model.User, error)
	// Login()
}

type RegisterUserDTO struct {
	Email    string
	Name     string
	Password string
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(ctx context.Context, input *RegisterUserDTO) (*model.User, error) {
	if input.Email == "" {
		return nil, errors.New("Email обязателен")
	}

	if !strings.Contains(input.Email, "@") || !strings.Contains(input.Email, ".") {
		return nil, errors.New("Невалидный формат email")
	}

	if input.Name == "" {
	}

	if input.Password == "" {
		return nil, errors.New("Пароль обязателен")
	}
	if len(input.Password) < 6 {
		return nil, errors.New("Пароль должен быть минимум на 6 символов")
	}

	// -----------

	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	createdUser, err := u.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

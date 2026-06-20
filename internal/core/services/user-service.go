package services

import (
	"context"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil {
		return err
	}

	user := &domain.User{
		Name: name,
		Password: string(hashedPassword),
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	return s.repo.GetUserByName(ctx, name)
}
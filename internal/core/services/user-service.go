package services

import (
	"context"
	"errors"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"
	"url-shortener-go/internal/adapters/outbound/utils"

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

func (s *UserService) Login(ctx context.Context, name, password string) (string, error) {
	user, err := s.repo.GetUserByName(ctx, name)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.Id, user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}
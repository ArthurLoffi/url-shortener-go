package services

import (
	"context"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string) error {
	user := &domain.User{
		Name: name,
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	return s.repo.GetUserByName(ctx, name)
}
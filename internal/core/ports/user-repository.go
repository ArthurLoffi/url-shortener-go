package ports

import (
	"context"
	"url-shortener-go/internal/core/domain"
)


type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetUserByName(ctx context.Context, name string) (*domain.User, error)
}
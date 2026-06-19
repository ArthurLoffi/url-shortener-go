package ports

import (
	"context"
	"url-shortener-go/internal/core/domain"
)

type UrlRepository interface {
	CreateUrl(ctx context.Context, url *domain.Url) error
	Redirect(ctx context.Context, code string) (*domain.Url, error)
	UpdateExpired(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.Url, error)
	GetByShortCode(ctx context.Context, code string) (*domain.Url, error)
	GetByUserID(ctx context.Context, userID uint) ([]domain.Url, error)
}
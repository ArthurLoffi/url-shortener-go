package ports

import (
	"context"
	"url-shortener-go/internal/core/domain"
)

type ClickRepository interface {
	Create(ctx context.Context, click *domain.Click) error
	GetByURLID(ctx context.Context, urlID uint) ([]domain.Click, error)
	CountByURLID(ctx context.Context, urlID uint) (int64, error)
}
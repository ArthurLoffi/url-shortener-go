package services

import (
	"context"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"
)

type ClickService struct {
	repo ports.ClickRepository
}

func NewClickService(repo ports.ClickRepository) *ClickService {
	return &ClickService{
		repo: repo,
	}
}

func (s *ClickService) Create(ctx context.Context, click *domain.Click) error {
	return s.repo.Create(ctx, click)
}

func (s *ClickService) GetByURLID(ctx context.Context, urlID uint) ([]domain.Click, error) {
	return s.repo.GetByURLID(ctx, urlID)
}

func (s *ClickService) CountByURLID(ctx context.Context, urlID uint) (int64, error) {
	return s.repo.CountByURLID(ctx, urlID)
}
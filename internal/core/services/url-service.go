package services

import (
	"context"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"
	"url-shortener-go/internal/core/utils"
)

type UrlService struct {
	repo ports.UrlRepository
}

func NewUrlService(repo ports.UrlRepository) *UrlService {
	return &UrlService{
		repo: repo,
	}
}

func (s *UrlService) CreateUrl(ctx context.Context, url *domain.Url) error {
	slug, err := utils.GenerateSlug()
	if err != nil {
		return err
	}

	url.ShortUrl = slug

	return s.repo.CreateUrl(ctx, url);
}

func (s *UrlService) Redirect(ctx context.Context, code string) (*domain.Url, error) {
	return s.repo.GetByShortCode(ctx, code)
}

func (s *UrlService) GetByID(ctx context.Context, id uint) (*domain.Url, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UrlService) GetByShortCode(ctx context.Context, code string) (*domain.Url, error) {
	return s.repo.GetByShortCode(ctx, code)
}

func (s *UrlService) GetByUserID(ctx context.Context, userID uint) ([]domain.Url, error){
	return s.repo.GetByUserID(ctx, userID)
}
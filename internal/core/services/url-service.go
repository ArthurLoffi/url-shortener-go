package services

import (
	"context"
	"fmt"
	"time"
	"url-shortener-go/internal/adapters/outbound/cache"
	"url-shortener-go/internal/adapters/outbound/utils"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/ports"
)

type UrlService struct {
	repo  ports.UrlRepository
	cache *cache.UrlCache
}

func NewUrlService(repo ports.UrlRepository, cache *cache.UrlCache) *UrlService {
	return &UrlService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UrlService) CreateUrl(ctx context.Context, url *domain.Url) error {
	slug, err := utils.GenerateSlug()
    if err != nil {
        return err
    }

    url.ShortUrl = slug

    return s.repo.CreateUrl(ctx, url)
}

func (s *UrlService) Redirect(ctx context.Context, code string) (*domain.Url, error) {
	url, err := s.repo.GetByShortCode(ctx, code)
	if err != nil {
		return nil, err
	}

    if time.Since(url.CreatedAt) > 24*time.Hour {
        s.repo.UpdateExpired(ctx, url.Id)
        return nil, fmt.Errorf("url expirada")
    }

    cached, err := s.cache.Get(ctx, code)
	if err == nil {
    	url.OriginalUrl = cached
 	   return url, nil
	}

	s.cache.Set(ctx, code, url.OriginalUrl)

	return url, nil
}

func (s *UrlService) GetByID(ctx context.Context, id uint) (*domain.Url, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UrlService) GetByShortCode(ctx context.Context, code string) (*domain.Url, error) {
	return s.repo.GetByShortCode(ctx, code)
}

func (s *UrlService) GetByUserID(ctx context.Context, userID uint) ([]domain.Url, error) {
	return s.repo.GetByUserID(ctx, userID)
}

func (s *UrlService) UpdateClickCount(ctx context.Context, urlID uint, count uint) error {
	return s.repo.UpdateClickCount(ctx, urlID, count)
}
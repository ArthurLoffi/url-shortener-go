package repository

import (
	"context"
	"url-shortener-go/internal/core/domain"

	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) CreateUrl(ctx context.Context, url *domain.Url) error {
	return r.db.WithContext(ctx).Create(url).Error
}

func (r *UrlRepository) Redirect(ctx context.Context, code string) (*domain.Url, error) {
    var url domain.Url
    err := r.db.WithContext(ctx).Where("short_url = ?", code).First(&url).Error
    return &url, err
}

func (r *UrlRepository) UpdateExpired(ctx context.Context, id uint) error {
	return r.db.Model(&domain.Url{}).Where("id = ?", id).Update("expired", true).Error
}

func (r *UrlRepository) GetByID(ctx context.Context, id uint) (*domain.Url, error) {
	var url domain.Url
	err := r.db.WithContext(ctx).First(&url, id).Error
	return &url, err 
}

func (r *UrlRepository) GetByShortCode(ctx context.Context, code string) (*domain.Url, error) {
	var url domain.Url
	err := r.db.WithContext(ctx).Where("short_url = ?", code).First(&url).Error
	return &url, err
}

func (r *UrlRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.Url, error) {
	var urls []domain.Url
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&urls).Error
	return urls, err
}
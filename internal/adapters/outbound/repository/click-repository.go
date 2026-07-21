package repository

import (
	"context"
	"url-shortener-go/internal/core/domain"

	"gorm.io/gorm"
)

type ClickRepository struct {
	db *gorm.DB
}

func NewClickRepository(db *gorm.DB) *ClickRepository {
	return &ClickRepository{
		db: db,
	}
}

func (r *ClickRepository) Create(ctx context.Context, click *domain.Click) error {
	return r.db.WithContext(ctx).Create(click).Error
}

func (r *ClickRepository) GetByURLID(ctx context.Context, urlID uint) ([]domain.Click, error) {
	var clicks []domain.Click
	err := r.db.WithContext(ctx).Where("url_id = ?", urlID).Find(&clicks).Error
	return clicks, err
}

func (r *ClickRepository) CountByURLID(ctx context.Context, urlID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Click{}).Where("url_id = ?", urlID).Count(&count).Error
	return count, err
}
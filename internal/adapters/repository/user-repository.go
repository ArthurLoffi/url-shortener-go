package repository

import (
	"context"
	"url-shortener-go/internal/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&user).Error
	return &user, err
}
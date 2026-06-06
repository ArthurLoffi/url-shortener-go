package domain

import "time"

type User struct {
	id uint `gorm:"primaryKey;autoincrement"`
	name string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdateAt time.Time
}
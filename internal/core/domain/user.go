package domain

import "time"

type User struct {
	Id uint `gorm:"primaryKey;autoincrement"`
	Name string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
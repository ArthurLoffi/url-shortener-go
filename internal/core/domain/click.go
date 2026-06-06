package domain

import "time"

type Click struct {
	id uint `gorm:"primaryKey;autoincrement"`
	urlid uint
	ipAdress string
	ClickedAt time.Time
}
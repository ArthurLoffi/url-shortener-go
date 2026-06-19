package domain

import "time"

type Url struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint `json:"user_id" gorm:"column:user_id"`
	OriginalUrl string  `json:"original_url" gorm:"column:original_url"`
	ShortUrl string `json:"short_url"`
	Clicks int64 `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
	Expired bool `json:"expired" gorm:"column:expired;default:false"`
}
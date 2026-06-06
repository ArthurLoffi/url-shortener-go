package domain

type Url struct {
	id uint `gorm:"primaryKey;autoincrement"`
	originalUrl string
	ShortUrl string
	Clicks int64
}
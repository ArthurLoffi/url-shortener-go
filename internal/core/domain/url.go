package domain

type Url struct {
	Id uint `json:"id" gorm:"primaryKey;autoincrement"`
	UserID uint `json:"user_id" gorm:"column:user_id"`
	OriginalUrl string  `json:"original_url" gorm:"column:original_url"`
	ShortUrl string `json:"short_url"`
	Clicks int64 `json:"clicks"`
}
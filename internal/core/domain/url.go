package domain

type Url struct {
	Id uint `json:"id" gorm:"primaryKey;autoincrement"`
	OriginalUrl string  `json:"original_url" gorm:"column:original_url"`
	ShortUrl string `json:"short_url"`
	Clicks int64 `json:"clicks"`
}
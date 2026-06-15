package domain

import "time"

type Click struct {
	Id uint `gorm:"primaryKey;autoincrement"`
	Urlid uint `gorm:"column:url_id"`
    IPAddress string `gorm:"column:ip_address"`
	ClickedAt time.Time
}
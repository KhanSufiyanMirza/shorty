package entity

import (
	"time"
)

type CustomShort struct {
	ID             uint      `gorm:"column:id;primary_key" json:"id"`
	CustomShortURL string    `gorm:"column:custom_short_url;unique_index;not null" json:"customShortUrl"`
	ActualURL      string    `gorm:"column:actual_url;not null" json:"actualUrl"`
	ExpiresAt      time.Time `gorm:"column:expires_at" json:"ttl"`
}

// TableName specifies the table name for the CustomShort model.
func (CustomShort) TableName() string {
	return "custom_shorts"
}

type RateLimit struct {
	IPAddress          string    `gorm:"column:ip_address;not null;unique_index" json:"ipAddress"`
	RemainingRateLimit int64     `gorm:"column:remaining_rate_limit" json:"remainingRateLimit"`
	ExpiresAt          time.Time `gorm:"column:expires_at" json:"ttl"`
}

func (RateLimit) TableName() string {
	return "rate_limits"
}

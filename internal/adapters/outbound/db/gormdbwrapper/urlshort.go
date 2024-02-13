package gormdbwrapper

import (
	"context"
	"errors"
	"hex/constants"
	"hex/models/entity"
	"hex/utils/logger"
	"time"

	"gorm.io/gorm"
)

//TODO: all cache values should be any type not string

// urlShorteningAdapter contain list of services required for urlShorteningAdapter
//
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . UrlShorteningDAO
type urlShorteningAdapter struct {
	db  *gorm.DB
	log logger.Logger
}

// DecrementAPIQuotaBalanceByOne implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) DecrementAPIQuotaBalanceByOne(ctx context.Context, IPAddress string) error {

	result := urlShorteningAdapter.db.Exec("UPDATE rate_limits SET remaining_rate_limit = remaining_rate_limit - 1 WHERE ip_address = ?", IPAddress)

	if result.RowsAffected == 0 {
		if result.Error != nil {
			return result.Error
		}
		return errors.New("quota exhausted")
	}
	return nil
}

// GetAPIQuotaBalanceBy implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetAPIQuotaBalanceBy(ctx context.Context, IPAddress string) (int64, error) {
	var limit entity.RateLimit
	result := urlShorteningAdapter.db.Where("ip_address = ?", IPAddress).First(&limit)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return 0, constants.ErrRecordNotFound // Return default limit if record doesn't exist
		}
		return 0, result.Error
	}

	return limit.RemainingRateLimit, nil
}

// GetAPIQuotaTTLBy implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetAPIQuotaTTLBy(ctx context.Context, IPAddress string) (time.Duration, error) {
	var limit entity.RateLimit
	result := urlShorteningAdapter.db.Where("ip_address = ?", IPAddress).First(&limit)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return time.Duration(0), nil // Return default TTL if record doesn't exist
		}
		return time.Duration(0), result.Error
	}

	if limit.ExpiresAt.After(time.Now()) {
		return time.Until(limit.ExpiresAt), nil
	} else {
		result := urlShorteningAdapter.db.Where("ip_address = ?", IPAddress).Delete(&limit)
		if result.Error != nil {
			return time.Duration(0), result.Error
		}

		return time.Duration(0), nil
	}

}

// SaveAPIQuota implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) SaveAPIQuota(ctx context.Context, IPAddress string, balance int64, ttl time.Duration) error {
	result := urlShorteningAdapter.db.Create(&entity.RateLimit{IPAddress: IPAddress, RemainingRateLimit: int64(balance), ExpiresAt: time.Now().Add(ttl)})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SaveCustomShort implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) SaveCustomShort(ctx context.Context, customShort string, url string, ttl time.Duration) error {
	result := urlShorteningAdapter.db.Create(&entity.CustomShort{CustomShortURL: customShort, ActualURL: url, ExpiresAt: time.Now().Add(ttl)})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUrlByCustomShort implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetUrlByCustomShort(ctx context.Context, customShort string) (string, error) {
	var short entity.CustomShort
	result := urlShorteningAdapter.db.Where("custom_short_url = ?", customShort).First(&short)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", constants.ErrRecordNotFound
		}
		return "", result.Error
	}

	if short.ExpiresAt.Before(time.Now()) {
		result := urlShorteningAdapter.db.Where("custom_short_url = ?", customShort).Delete(&short)
		if result.Error != nil {
			return "", result.Error
		}
		return "", errors.New("short URL expired")
	}

	return short.ActualURL, nil
}

// NewAdapter creates a new Adapter
func NewUrlShorteningAdapter(da *DBWrapper) *urlShorteningAdapter {

	return &urlShorteningAdapter{db: da.db, log: da.log}
}

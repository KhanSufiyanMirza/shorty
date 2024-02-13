package cachestore

import (
	"context"
	"errors"
	"hex/constants"
	"hex/internal/ports/outbound"
	"strconv"
	"time"
)

type urlShorteningAdapter struct {
	db outbound.IRedis
}

// DecrementAPIQuotaBalanceByOne implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) DecrementAPIQuotaBalanceByOne(ctx context.Context, IPAddress string) error {
	return urlShorteningAdapter.db.Decr(ctx, IPAddress)
}

// GetAPIQuotaBalanceBy implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetAPIQuotaBalanceBy(ctx context.Context, IPAddress string) (int64, error) {
	val, err := urlShorteningAdapter.db.Get(ctx, IPAddress)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return 0, constants.ErrRecordNotFound
		}
		return 0, err
	}
	remainingRateLimit, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return int64(remainingRateLimit), nil
}

// GetAPIQuotaTTLBy implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetAPIQuotaTTLBy(ctx context.Context, IPAddress string) (time.Duration, error) {
	ttl, err := urlShorteningAdapter.db.GetTTL(ctx, IPAddress)
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

// GetUrlByCustomShort implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) GetUrlByCustomShort(ctx context.Context, customShort string) (string, error) {
	url, err := urlShorteningAdapter.db.Get(ctx, customShort)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return "", constants.ErrRecordNotFound
		}
		return "", err
	}
	return url, nil
}

// SaveAPIQuota implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) SaveAPIQuota(ctx context.Context, IPAddress string, balance int64, ttl time.Duration) error {
	return urlShorteningAdapter.db.SetWithTTL(ctx, IPAddress, balance, ttl)
}

// SaveCustomShort implements outbound.UrlShortenerDAO.
func (urlShorteningAdapter *urlShorteningAdapter) SaveCustomShort(ctx context.Context, customShort string, url string, ttl time.Duration) error {
	return urlShorteningAdapter.db.SetWithTTL(ctx, customShort, url, ttl)
}

// NewAdapter creates a new Adapter
func NewUrlShorteningAdapter(db outbound.IRedis) outbound.UrlShortenerDAO {

	return &urlShorteningAdapter{db: db}
}

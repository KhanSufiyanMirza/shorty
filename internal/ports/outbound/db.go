package outbound

import (
	"context"
	"time"
)

type DbPort interface {
	GetCacheStore() IRedis
	GetUrlShortenerDAO() UrlShortenerDAO
}

type DbOps struct {
	CacheStore      IRedis
	UrlShortenerDAO UrlShortenerDAO
}

type DbOpsFunc func(*DbOps)

type DefaultDbPort interface {
	// GetDefault returns an instance of the outbound.DbOps
	GetDefault() DbOps
}

type UrlShortenerDAO interface {
	GetAPIQuotaBalanceBy(ctx context.Context, IPAddress string) (int64, error)
	GetAPIQuotaTTLBy(ctx context.Context, IPAddress string) (time.Duration, error)
	SaveAPIQuota(ctx context.Context, IPAddress string, balance int64, ttl time.Duration) error
	DecrementAPIQuotaBalanceByOne(ctx context.Context, IPAddress string) error
	GetUrlByCustomShort(ctx context.Context, customShort string) (string, error)
	SaveCustomShort(ctx context.Context, customShort string, url string, ttl time.Duration) error
}

package db

import (
	"hex/internal/ports/outbound"
	// blank import for mysql driver

	_ "github.com/go-sql-driver/mysql"
)

// Adapter implements the DbPort interface
type Adapter struct {
	outbound.DbOps
}

// WithCacheStore returns outbound.DbOpsFunc.
// sets the CacheStore to the provided  IRedis instance.
func WithCacheStore(r outbound.IRedis) outbound.DbOpsFunc {
	return func(o *outbound.DbOps) {
		o.CacheStore = r
	}
}

// NewAdapter creates a new Adapter
func NewAdapter(d outbound.DefaultDbPort, funcs ...outbound.DbOpsFunc) *Adapter {
	ops := d.GetDefault()
	for _, fn := range funcs {
		fn(&ops)
	}

	return &Adapter{DbOps: ops}
}

// GetCacheStore is a method of the Adapter type that returns the CacheStore field.
// It provides access to the outbound.CacheStore instance associated with the Adapter.
func (da Adapter) GetCacheStore() outbound.IRedis {
	return da.CacheStore
}

func (da Adapter) GetUrlShortenerDAO() outbound.UrlShortenerDAO {
	return da.UrlShortenerDAO
}

package outbound

import (
	"context"
	"time"
)

type Initializer interface {
	Init() error
	Close() error
}

type IRedis interface {
	Initializer
	Context() context.Context
	CoreOps
	HashOps
	ListOps
	CounterOps
	PubSubOps
}

// CoreOps contains functions which allows to perform get, set, delete and append operations with redis db.
// With different get, set functions on basis of time of expiry of key.
type CoreOps interface {
	Ping(ctx context.Context) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
	GetTTL(ctx context.Context, key string) (time.Duration, error)
	SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error)
	GetAllWithPrefix(ctx context.Context, prefix string) ([]string, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) bool
	Append(ctx context.Context, key, appendValue string) error
	ConfigSet(ctx context.Context, parameter, value string) (string, error)
}

// CounterOps contains functions which perform increment/decrement operations on value of key in redis
type CounterOps interface {
	Incr(ctx context.Context, key string) error
	IncrBy(ctx context.Context, key string, value int64) error
	IncrByFloat(ctx context.Context, key string, value float64) error
	Decr(ctx context.Context, key string) error
	DecrBy(ctx context.Context, key string, value int64) error
}

// HashOps contains hset, hget functions which operate on fields and value in hash of redis
// It also contains set and get functions for values of field in a hash associated with the key
type HashOps interface {
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (interface{}, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HMSet(ctx context.Context, key string, values map[string]interface{}) error
	// HMGet(ctx context.Context, key string, fields ...string) error
	HMGet(ctx context.Context, key string, model interface{}, fields ...string) error
}

// ListOps contains functions for working with lists associated with the key in redis.
// Push, pop, indexing and ranging of elements in a list are key operations of functions
type ListOps interface {
	RPush(ctx context.Context, key string, values ...interface{}) error
	RPop(ctx context.Context, key string) (string, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) error
	LIndex(ctx context.Context, key string, index int64) (string, error)
	LRange(ctx context.Context, key string, start int64, end int64) ([]string, error)
}

// PubSubOps contains PSubsribe functionality of redis, which subscribes the client to the given patterns
type PubSubOps interface {
	PSubscribe(ctx context.Context, isReceived chan struct{}, values ...string)
}

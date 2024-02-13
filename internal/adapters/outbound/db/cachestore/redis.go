package cachestore

import (
	"context"
	"hex/config"
	"hex/internal/ports/outbound"
	"hex/utils/logger"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// type Config struct {
// 	Adders   map[string]string
// 	Password string
// 	DB       int
// }

// func NewConfig(hosts []string, pass string, dbNum int) Config {
// 	adders := hostMap(hosts)

// 	return Config{
// 		Adders:   adders,
// 		Password: pass,
// 		DB:       dbNum,
// 	}
// }

// hostMap to Converts list of hosts to map
func hostMap(hosts []string) map[string]string {
	hostMap := make(map[string]string)
	for i, host := range hosts {
		hostMap["shard"+strconv.Itoa(i)] = host
	}

	return hostMap
}

// RedisDB Ring internally exclude dead hosts from the ring and
// does consistent hashing among the live hosts.
type RedisDB struct {
	Log    logger.Logger
	Config config.CacheConfig
	Client *redis.Ring
}

// NewRedisDB creates a new instance of RedisDB by initializing it with the provided logger and cache configuration.
// Firstly, It initializes the RedisDB instance, And log's a fatal error if error occurs.
// Returns a pointer to the initialized RedisDB instance.
func NewRedisDB(log logger.Logger, cfg config.CacheConfig) *RedisDB {

	rdb := RedisDB{
		Log:    log,
		Config: cfg,
		Client: nil,
	}
	if err := rdb.Init(); err != nil {
		log.Fatal("CacheStoreModule", "error while init redis", err)
		// log.Fatal("RedisDb", "Error while Redis DB init", err)
	}

	return &rdb

}

// WithCacheStore is a function that modifies DbOps configuration.
func (redisDB *RedisDB) WithCacheStore() outbound.DbOpsFunc {
	return func(o *outbound.DbOps) {
		o.CacheStore = redisDB
	}
}

// WithCacheStore is a function that modifies DbOps configuration.
func (redisDB *RedisDB) WithUrlShortener() outbound.DbOpsFunc {
	return func(o *outbound.DbOps) {
		o.UrlShortenerDAO = NewUrlShorteningAdapter(redisDB)
	}
}

// Init initializes the RedisDB instance by creating a Redis client by using the provided configuration.
// Firstly, Creating a Redis client using the configuration options.
// And then Getting the context for the Redis client and Ping the Redis server.
// Now, Checking and returning the error occured during the ping operation.
func (rdb *RedisDB) Init() error {
	rdb.Client = redis.NewRing(&redis.RingOptions{ //nolint:exhaustruct
		Addrs:    rdb.Config.Adders,
		Password: rdb.Config.Password,
		DB:       rdb.Config.DB,
	})
	ctx := rdb.Client.Context()
	if _, err := rdb.Ping(ctx); err != nil {
		return err
	}

	return nil

}

// Context used to retrieve and return the context used by the Redis client.
func (rdb *RedisDB) Context() context.Context {
	return rdb.Client.Context()
}

// Close function used to close the redis db connection.
// Returns an error if there's any issue occured while closing the connection.
func (rdb *RedisDB) Close() error {
	if err := rdb.Client.Close(); err != nil {
		return errors.Wrap(err, "Error While Closing Connection")
	}

	return nil
}

// Ping sends a PING command to the Redis server to check connectivity.
// Execute the Ping command on the Redis client with the given context and retrieves the PONG response.
// If the PONG response is received, It indicates that successfully connected to redis server.
func (rdb *RedisDB) Ping(ctx context.Context) (bool, error) {
	pong, err := rdb.Client.WithContext(ctx).Ping(ctx).Result()
	if err != nil {
		return false, errors.Wrap(err, "Error while Ping")
	}

	return pong == "PONG", nil
}

// Get function is used to get the value for the key from the redis server.
// Firstly, Checking if the key is empty. If its empty, returning an error.
// Now, Retrieving the value with the given key from the Redis server using the context.
// If the key is not found in Redis, return an error.
// Finally, Return the retrieved value as a string.
func (rdb *RedisDB) Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	val, err := rdb.Client.WithContext(ctx).Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}

		return "", errors.Wrap(err, "Error While Get")
	}

	return val, nil
}

// Set function used to set the value with the provided key in the Redis server, with default 0 expiration.
// Firstly, Checking if the key is empty. If its empty, returning an error.
// Now, Executing the SET command with the given key and value on the Redis server.
// If an error occurs during the SET operation, it returns the error.
func (rdb *RedisDB) Set(ctx context.Context, key string, value interface{}) error {
	if key == "" {
		return ErrEmptyKey
	}
	err := rdb.Client.WithContext(ctx).Set(ctx, key, value, 0).Err()
	if err != nil {
		return errors.Wrap(err, "Error While Set")
	}

	return nil
}

// GetTTL Returns remaining time to expire a key in the redis server.
// Firstly, It Retrieves the TTL duration for the given key from the Redis server.
// If the key is not found in Redis, it returns 0 duration and an error indicating key not found.
// Finally, returns the TTL duration for the key.
func (rdb *RedisDB) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := rdb.Client.WithContext(ctx).TTL(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, ErrKeyNotFound
		}

		return 0, errors.Wrap(err, "Error while GetTTL")
	}

	return ttl, nil
}

// SetWithTTL function used to set key value with TTL in the redis server.
// Firstly, Checking if the key is empty. If its empty, returning an error.
// Now, running the SET command with the provided key, value, and TTL on the Redis server.
// If an error occurs during the SET operation with TTL, returns the error.
func (rdb *RedisDB) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}
	err := rdb.Client.WithContext(ctx).Set(ctx, key, value, ttl).Err()
	if err != nil {
		return errors.Wrap(err, "Error While Setting a key value with TTL")
	}

	return nil
}

// SetNX sets the value with the key in the Redis server with a TTL, only if the key does not already exist.
// Firstly, Checking if the key is empty. If its empty, returning an error.
// Now, running the SETNX command with the provided key, value, and TTL on the Redis server.
// If an error occurs during the SETNX operation, returns an error.
func (rdb *RedisDB) SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	if key == "" {
		return false, ErrEmptyKey
	}
	ok, err := rdb.Client.WithContext(ctx).SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return false, errors.Wrap(err, "Error while acquiring the lock")
	}

	return ok, nil
}

// GetAllWithPrefix retrieves all the prefixes from a Redis database.
// Firstly, Iterating through the keys matching the prefix until all are retrieved.
// Now, Creating a temporary storage for keys retrieved in each scan iteration.
// And then Performing a SCAN operation in Redis with the given prefix pattern.
// It Return's an error, if there's an error during scanning.
// Also it appends the retrieved keys to the keys slice.
// Now, Checking if the cursor returned is 0, indicating that no more keys to retrieve.
// Finally, returns the keys that are matching the prefix.
func (rdb *RedisDB) GetAllWithPrefix(ctx context.Context, prefix string) ([]string, error) {
	var keys []string
	var cursor uint64
	var err error

	for {
		var scanKeys []string
		scanKeys, cursor, err = rdb.Client.Scan(ctx, cursor, prefix+"*", 0).Result()
		if err != nil {
			return nil, errors.Wrap(err, "Error While GetAllWithPrefix")
		}

		keys = append(keys, scanKeys...)

		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

// Delete to delete key from the redis database.
// It Uses the Redis client to delete the key in the given context.
func (rdb *RedisDB) Delete(ctx context.Context, key string) error {
	err := rdb.Client.WithContext(ctx).Del(ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, "Error While Deleting a key")
	}

	return nil
}

// Exists checks if a key exists in the Redis database.
func (rdb *RedisDB) Exists(ctx context.Context, key string) bool {
	val, _ := rdb.Client.WithContext(ctx).Exists(ctx, key).Result()

	return val != 0
}

// Append to add value + appendValue for the key (val type str only)
// Append method used to appends a value to a Redis key.
// Using the Redis client, perform the append operation on the specified key with the given value.
// The operation is executed within the provided context.
func (rdb *RedisDB) Append(ctx context.Context, key, appendValue string) error {
	_, err := rdb.Client.WithContext(ctx).Append(ctx, key, appendValue).Result()
	if err != nil {
		return errors.Wrap(err, "Error While Append")
	}

	return nil
}

// Incr increments the value of a key by 1 in Redis.
func (rdb *RedisDB) Incr(ctx context.Context, key string) error {
	err := rdb.Client.WithContext(ctx).Incr(ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, "Error While Incr")
	}

	return nil
}

// IncrBy increments the value of a key in Redis by a given amount.
func (rdb *RedisDB) IncrBy(ctx context.Context, key string, value int64) error {
	err := rdb.Client.WithContext(ctx).IncrBy(ctx, key, value).Err()

	if err != nil {
		return errors.Wrap(err, "Error While IncrBy")
	}

	return nil
}

// IncrByFloat increments the value of a key in Redis by a given floating-point value.
func (rdb *RedisDB) IncrByFloat(ctx context.Context, key string, value float64) error {
	err := rdb.Client.WithContext(ctx).IncrByFloat(ctx, key, value).Err()

	if err != nil {
		return errors.Wrap(err, "Error While IncrByFloat")
	}

	return nil
}

// Decr decrements the value of a key in Redis.
func (rdb *RedisDB) Decr(ctx context.Context, key string) error {
	err := rdb.Client.WithContext(ctx).Decr(ctx, key).Err()

	if err != nil {
		return errors.Wrap(err, "Error While Decr")
	}

	return nil
}

// DecrBy decrements the value of a key in Redis by a given amount.
func (rdb *RedisDB) DecrBy(ctx context.Context, key string, value int64) error {
	err := rdb.Client.WithContext(ctx).DecrBy(ctx, key, value).Err()

	if err != nil {
		return errors.Wrap(err, "Error While DecrBy")
	}

	return nil
}

// HSet sets fields and value in a hash in Redis.
func (rdb *RedisDB) HSet(ctx context.Context, key string, field string, value interface{}) error {
	if key == "" {
		return ErrEmptyKey
	}
	if field == "" {
		return ErrEmptyFields
	}
	err := rdb.Client.WithContext(ctx).HSet(ctx, key, field, value).Err()
	if err != nil {
		return errors.Wrap(err, "Error While HSet")
	}

	return nil
}

// HGet to get fields and value in a hash in db.
func (rdb *RedisDB) HGet(ctx context.Context, key string, field string) (interface{}, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	if field == "" {
		return nil, ErrEmptyFields
	}
	res, err := rdb.Client.WithContext(ctx).HGet(ctx, key, field).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrFieldNotFound
		}

		return nil, errors.Wrap(err, "Error While HGet")
	}

	return res, nil
}

// HGetAll returns all fields and values of a hash in Redis.
func (rdb *RedisDB) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	res, err := rdb.Client.WithContext(ctx).HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Error While HGetAll")
	}

	return res, nil
}

// HMSet sets the hash field to the value in the hash stored at key.
func (rdb *RedisDB) HMSet(ctx context.Context, key string, values map[string]interface{}) error {
	if key == "" {
		return ErrEmptyKey
	}
	if len(values) == 0 {
		return ErrEmptyValues
	}
	err := rdb.Client.WithContext(ctx).HMSet(ctx, key, values).Err()
	if err != nil {
		return errors.Wrap(err, "Error while HMSet")
	}

	return nil
}

// HMGet returns the value associated with the field in the hash stored at key.
func (rdb *RedisDB) HMGet(ctx context.Context, key string, model interface{}, fields ...string) error {
	if key == "" {
		return ErrEmptyKey
	}
	if len(fields) == 0 {
		return ErrEmptyFields
	}
	err := rdb.Client.WithContext(ctx).HMGet(ctx, key, fields...).Scan(&model)
	if err != nil {
		return errors.Wrap(err, "Error while HMGet")
	}

	return nil
}

// RPush Push back to the list associated with key.
func (rdb *RedisDB) RPush(ctx context.Context, key string, values ...interface{}) error {
	err := rdb.Client.WithContext(ctx).RPush(ctx, key, values...).Err()
	if err != nil {
		return errors.Wrap(err, "Error While RPush")
	}

	return nil
}

// RPop to Pop back to the list associated with key.
func (rdb *RedisDB) RPop(ctx context.Context, key string) (string, error) {
	val, err := rdb.Client.WithContext(ctx).RPop(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}

		return "", errors.Wrap(err, "Error While RPop")
	}

	return val, nil
}

// LSet to Insert value at an index of list associated with key
func (rdb *RedisDB) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	err := rdb.Client.WithContext(ctx).LSet(ctx, key, index, value).Err()
	if err != nil {
		return errors.Wrap(err, "Error While LSet")
	}

	return nil
}

// LIndex Returns value at index of list associated with key.
func (rdb *RedisDB) LIndex(ctx context.Context, key string, index int64) (string, error) {
	val, err := rdb.Client.WithContext(ctx).LIndex(ctx, key, index).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrIndexOutOfRange
		}

		return "", errors.Wrap(err, "Error While LIndex")
	}

	return val, nil
}

// LRange returns a range of elements from a list.
func (rdb *RedisDB) LRange(ctx context.Context, key string, start int64, end int64) ([]string, error) {
	val, err := rdb.Client.LRange(ctx, key, start, end).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Error While LRange")
	}

	return val, nil
}

// PSubscribe subscribes to Redis channels and listens for incoming messages.
func (rdb *RedisDB) PSubscribe(ctx context.Context, isReceived chan struct{}, values ...string) {

	pubSub := rdb.Client.WithContext(ctx).PSubscribe(ctx, values...)

	defer func() {
		pubSub.Close()
		close(isReceived)
	}()

	for range pubSub.Channel() {
		isReceived <- struct{}{}

	}

}

// ConfigSet sets a configuration parameter in Redis.
func (rdb *RedisDB) ConfigSet(ctx context.Context, parameter, value string) (string, error) {
	return rdb.Client.ConfigSet(ctx, parameter, value).Result()
}

var ErrEmptyKey = errors.New("key is nil")
var ErrEmptyValues = errors.New("values are nil")
var ErrEmptyFields = errors.New("field is nil")
var ErrKeyNotFound = errors.New("key not found")
var ErrFieldNotFound = errors.New("field not present")
var ErrIndexOutOfRange = errors.New("index out of range")

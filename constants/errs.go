package constants

import "errors"

// General error messages
var (
	ErrRecordAlreadyExisting = errors.New("record already exist")
	ErrRecordNotFound        = errors.New("record not found")
	ErrCacheAlreadyExist     = errors.New("cache already exists")
	ErrCacheNotFound         = errors.New("cache not found")
	ErrRateLimitExceeded     = errors.New("rate limit exceeded")
	ErrInvalidUrl            = errors.New("invalid url")
	ErrDomainUrl             = errors.New("remove domain error")
)

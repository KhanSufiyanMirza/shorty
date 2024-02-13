package dto

import "time"

type CustomShortRequest struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type CustomShortResponse struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int64         `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

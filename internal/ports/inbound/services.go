package inbound

import (
	"context"
	"hex/models/dto"
)

// TODO: need to refactor
// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	GetUrlShortingService() UrlShortingService
}

// IndentService contains fetch, save, update functions for operating with indents
type UrlShortingService interface {
	CreateShorty(ctx context.Context, req *dto.CustomShortRequest, ipAddress string) (dto.CustomShortResponse, error)
	ResolveShorty(ctx context.Context, customUrl string) (string, error)
}

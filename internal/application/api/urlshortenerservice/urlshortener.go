package urlshortenerservice

import (
	"context"
	"errors"
	"fmt"
	"hex/config"
	"hex/constants"
	"hex/internal/ports/outbound"
	"hex/models/dto"
	"hex/utils/helpers"
	"hex/utils/logger"
	"time"

	"github.com/google/uuid"
)

// Service represents list of services required for Service
type Service struct {
	db     outbound.DbPort
	log    logger.Logger
	config config.Config
}

// CreateShorty implements inbound.UrlShortingService.
func (service *Service) CreateShorty(ctx context.Context, req *dto.CustomShortRequest, ipAddress string) (dto.CustomShortResponse, error) {

	leftApiQuotaBalance, err := service.db.GetUrlShortenerDAO().GetAPIQuotaBalanceBy(ctx, ipAddress)
	if err != nil {
		if !errors.Is(err, constants.ErrRecordNotFound) {
			return dto.CustomShortResponse{}, err
		}
		err = service.db.GetUrlShortenerDAO().SaveAPIQuota(ctx, ipAddress, service.config.APIQuota, service.config.APIQuotaTTL)
		if err != nil {
			return dto.CustomShortResponse{}, err
		}
	} else {
		if leftApiQuotaBalance <= 0 {
			ttl, err := service.db.GetUrlShortenerDAO().GetAPIQuotaTTLBy(ctx, ipAddress)
			if err != nil {
				return dto.CustomShortResponse{}, err
			}
			return dto.CustomShortResponse{}, fmt.Errorf("%w rate_limit_reset: %v", constants.ErrRateLimitExceeded, ttl/time.Nanosecond/time.Minute)
		}
	}

	if helpers.RemoveDomainError(req.URL, service.config.DomainUrl) {
		return dto.CustomShortResponse{}, constants.ErrDomainUrl
	}
	req.URL = helpers.EnforceHTTP(req.URL)
	var customShortUrl string
	if req.CustomShort == "" {
		customShortUrl = uuid.New().String()[:6]
	} else {
		customShortUrl = req.CustomShort
	}

	url, err := service.db.GetUrlShortenerDAO().GetUrlByCustomShort(ctx, customShortUrl)
	if err != nil && !errors.Is(err, constants.ErrRecordNotFound) {
		return dto.CustomShortResponse{}, err
	}
	if url != "" {
		return dto.CustomShortResponse{}, constants.ErrRecordAlreadyExisting
	}
	if req.Expiry == 0 {
		req.Expiry = 24 // default expiry of 24 hours
	}
	err = service.db.GetUrlShortenerDAO().SaveCustomShort(ctx, customShortUrl, req.URL, req.Expiry*time.Hour)
	if err != nil {
		return dto.CustomShortResponse{}, err
	}
	err = service.db.GetUrlShortenerDAO().DecrementAPIQuotaBalanceByOne(ctx, ipAddress)
	if err != nil {
		return dto.CustomShortResponse{}, err
	}
	val, err := service.db.GetUrlShortenerDAO().GetAPIQuotaBalanceBy(ctx, ipAddress)
	if err != nil {
		return dto.CustomShortResponse{}, err
	}
	ttl, err := service.db.GetUrlShortenerDAO().GetAPIQuotaTTLBy(ctx, ipAddress)
	if err != nil {
		return dto.CustomShortResponse{}, err
	}
	var resp = dto.CustomShortResponse{
		XRateRemaining:  val,
		XRateLimitReset: ttl / time.Nanosecond / time.Minute,
		URL:             req.URL,
		CustomShort:     service.config.DomainUrl + "/" + customShortUrl,
	}

	return resp, nil

}

// ResolveShorty implements inbound.UrlShortingService.
func (service *Service) ResolveShorty(ctx context.Context, customUrl string) (string, error) {
	url, err := service.db.GetUrlShortenerDAO().GetUrlByCustomShort(ctx, customUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}

// New creates and initializes a new instance of the Service struct.
// Returns a pointer to the initialized Service.
func New(
	db outbound.DbPort,
	log logger.Logger,
	config config.Config,

) *Service {
	return &Service{
		db:     db,
		log:    log,
		config: config,
	}
}

package api

import (
	"context"
	"hex/config"
	"hex/internal/application/api/urlshortenerservice"
	"hex/internal/ports/inbound"
	"hex/internal/ports/outbound"
	"hex/utils/logger"
)

// Application implements the APIPort interface
type Application struct {
	db                 outbound.DbPort
	urlShortingService inbound.UrlShortingService
}

// NewApplication creates a new Application
func NewApplication(ctx context.Context, config config.Config, database outbound.DbPort, logger logger.Logger) *Application {
	return &Application{
		db:                 database,
		urlShortingService: urlshortenerservice.New(database, logger, config),
	}
}

func (app *Application) GetUrlShortingService() inbound.UrlShortingService {
	return app.urlShortingService
}

package main

import (
	"context"
	"hex/cmd/http/server"
	"hex/config"
	"hex/internal/adapters/outbound/db"
	"hex/internal/adapters/outbound/db/cachestore"
	"hex/internal/adapters/outbound/db/gormdbwrapper"
	"hex/internal/application/api"
	"hex/utils/logger"
)

func main() {
	ctx := context.Background()
	config := config.NewAppConfig()
	config.Validate(false, true, false, false)
	logger := logger.Init(ctx, logger.DebugLevel, config.AppName)
	dbWrapper, err := gormdbwrapper.NewDBWrapper(logger, config.ProvideRDBConfig(), config.ProvideCacheConfig())
	if err != nil {
		logger.Fatal("Main", "error while Initializing DB %v", err)
	}
	defer dbWrapper.Close()

	cachestore := cachestore.NewRedisDB(logger, config.ProvideCacheConfig())
	defer dbWrapper.Close()

	dbAdapter := db.NewAdapter(dbWrapper, cachestore.WithUrlShortener())

	applicationAPI := api.NewApplication(ctx, config, dbAdapter, logger)

	server.New(&config, applicationAPI, logger, dbAdapter).Run()
}

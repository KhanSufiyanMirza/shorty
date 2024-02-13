package config

import "hex/models/entity"

// ProvideRDBConfig used to generate RDBConfigOptions using AppConfig and a list of serverEntities.
// Firstly, passing a list of serverEntities representing various models.
// And returns RDBConfigOptions using the provided AppConfig and serverEntities.
func (appConfig AppConfig) ProvideRDBConfig() RDBConfigOptions {
	serverEntities := []interface{}{
		entity.RateLimit{},
		entity.CustomShort{},
	}

	return newRDBConfigOptions(appConfig.RDBConfig.Host,
		appConfig.RDBConfig.User,
		appConfig.RDBConfig.Passwd,
		appConfig.RDBConfig.DBName,
		appConfig.AppName,
		serverEntities...,
	)

}

// ProvideCacheConfig returns a CacheConfig based on the AppConfig.
// Firstly, Creating a new CacheConfig instance with the KeyDB host address
// and password from the AppConfig, And Setting expiration time to 0 for cache items.
func (appConfig AppConfig) ProvideCacheConfig() CacheConfig {
	return NewCacheConfig([]string{appConfig.KeyDBConfig.Host},
		appConfig.KeyDBConfig.Passwd, 0)
}

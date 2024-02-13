package gormdbwrapper

import (
	"fmt"
	"hex/config"
	"hex/internal/adapters/outbound/db/cachestore"
	"hex/internal/ports/outbound"
	"hex/utils/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const (
	// cockroach related
	dataSourceNoPasswordURIFmt = "postgresql://%s@%s/%s?sslmode=disable&application_name=$%s"
	dataSourceURIFmt           = "postgresql://%s:%s@%s/%s?sslmode=disable&application_name=$%s"

	// maria related

	mariaConnection = "%s:%s@tcp(%s)/%s?parseTime=true&tls=false&multiStatements=true&interpolateParams=true"

	maxIdleConnection     = 250
	maxOpenConnection     = 500
	maxConnectionLifetime = 5 * time.Minute
	// defaultDb             = "mysql"
	defaultDb = "postgres"
)

// TODO: move to outbound ports

// BadgerWrapper is a wrapper around Badger DB.
type DBWrapper struct {
	log         logger.Logger
	config      config.RDBConfigOptions
	db          *gorm.DB
	ErrorCode   config.ErrorCode
	cacheStore  outbound.IRedis
	cacheConfig config.CacheConfig
}

// NewDBWrapper creates a new instances of gorm db and cache db (redis).
func NewDBWrapper(log logger.Logger, config config.RDBConfigOptions, cacheConfig config.CacheConfig) (*DBWrapper, error) {

	if config.DBName == "" || config.Host == "" {
		return nil, fmt.Errorf("invalid config for host:%v", config.Host)
	}

	DBWrapper := &DBWrapper{
		log:         log,
		config:      config,
		cacheConfig: cacheConfig,
		db:          &gorm.DB{},

		cacheStore: cachestore.NewRedisDB(log, cacheConfig),
		// ErrorCode: pgerror.PgError{},

	}
	err := DBWrapper.init()
	if err != nil {
		return nil, fmt.Errorf("failed to open Badger DB: %w", err)
	}

	return DBWrapper, nil
}

// Close closes the Badger DB.
func (bd *DBWrapper) Close() error {
	if bd.db != nil {

		db, err := bd.db.DB()
		if err != nil {
			return fmt.Errorf("failed to close RDB: %w", err)
		}

		db.Close()
	}

	if bd.cacheStore != nil {
		err := bd.cacheStore.Close()
		if err != nil {
			return fmt.Errorf("failed to close cachestore: %w", err)
		}

	}

	return nil
}

// GetDefault returns an instance of the outbound.DbOps
func (bd *DBWrapper) GetDefault() outbound.DbOps {
	return outbound.DbOps{

		UrlShortenerDAO: NewUrlShorteningAdapter(bd),

		CacheStore: bd.cacheStore,
	}
}

// init initializes the database connection by selecting the appropriate database type
// (CockroachDB or MariaDB), establishing the connection pool, and performing auto-migration
// for the specified models. It sets the initialized GORM database instance in the DBWrapper.
// The method returns an error if any of these steps fail.
func (g *DBWrapper) init() error {
	var (
		db  *gorm.DB
		err error
	)
	switch defaultDb {
	case "postgres":
		db, err = g.NewCockroachDBConnectionPool()
		if err != nil {
			err = fmt.Errorf("cockroachdb: %v", err)
		}
	default:
		db, err = g.NewMariaDBConnectionPool()
		if err != nil {
			err = fmt.Errorf("mariadb: %v", err)
		}
	}

	if err != nil {
		g.log.Fatal("DbClientModule", "db connection pool failed", err)
	}

	err = db.AutoMigrate(g.config.Models...)
	if err != nil {
		g.log.Fatal("DbClientModule", "migration for models failed", err)
	}

	g.db = db
	return err
}

// NewCockroachDBConnectionPool establishes a connection pool to a CockroachDB database
// using the configuration provided in the DBWrapper instance. It returns a GORM database.
// Retry the connection setup if an error occurs, up to the specified maximum retries.
// Checks if database exists, if not exist then creates new database.
// Sets the maximum idle connections, maximum open connections, and connection lifetime.
// It returns a GORM database instance and an error,
// if any, encountered during the connection setup.
func (g DBWrapper) NewCockroachDBConnectionPool() (*gorm.DB, error) {
	var dsn string
	if g.config.Password == "" {
		dsn = fmt.Sprintf(dataSourceNoPasswordURIFmt, g.config.Username, g.config.Host, g.config.DBName, g.config.AppName)
	} else {
		dsn = fmt.Sprintf(dataSourceURIFmt, g.config.Username, g.config.Password, g.config.Host, g.config.DBName, g.config.AppName)
	}

	dblogger := logger.GetDBLogger()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dblogger.LogMode(gormlogger.Silent),
	})
	retryCount := 0
	maxRetry := 3
	for err != nil {
		dblogger := logger.GetDBLogger()
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: dblogger.LogMode(gormlogger.Silent),
		})
		if err != nil {
			retryCount++
			if retryCount > maxRetry {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()

	rawQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", g.config.DBName)
	dbCreateError := db.Raw(rawQuery).Error
	if dbCreateError != nil {
		return nil, dbCreateError
	}
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConnection)
	sqlDB.SetMaxOpenConns(maxOpenConnection)
	sqlDB.SetConnMaxLifetime(maxConnectionLifetime)
	return db, nil
}

// NewMariaDBConnectionPool establishes a connection pool to Maria Database
// using the configuration provided in the DBWrapper instance. It returns a GORM database.
// Retry the connection setup if an error occurs, up to the specified maximum retries.
// Checks if database exists, if not exist then creates new database.
// Sets the maximum idle connections, maximum open connections, and connection lifetime.
// It returns a GORM database instance and an error,
// if any, encountered during the connection setup.
func (g DBWrapper) NewMariaDBConnectionPool() (*gorm.DB, error) {

	dsn := fmt.Sprintf(mariaConnection, g.config.Username, g.config.Password, g.config.Host, g.config.DBName)

	dblogger := logger.GetDBLogger()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dblogger.LogMode(gormlogger.Silent),
	})
	retryCount := 0
	maxRetry := 3
	for err != nil {
		dblogger := logger.GetDBLogger()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: dblogger.LogMode(gormlogger.Silent),
		})
		if err != nil {
			retryCount++
			if retryCount > maxRetry {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	//check if database exists if not exist then create it

	rawQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", g.config.DBName)
	dbCreateError := db.Raw(rawQuery).Error
	if dbCreateError != nil {
		return nil, dbCreateError
	}
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConnection)
	sqlDB.SetMaxOpenConns(maxOpenConnection)
	sqlDB.SetConnMaxLifetime(maxConnectionLifetime)
	return db, nil
}

// WithCacheStore returns outbound.DbOpsFunc.
// Sets the CacheStore with a new instance of the NewRedisDB using the provided
// logger and cache configuration from the DBWrapper.
func (bd *DBWrapper) WithCacheStore() outbound.DbOpsFunc {
	return func(o *outbound.DbOps) {
		o.CacheStore = cachestore.NewRedisDB(bd.log, bd.cacheConfig)
	}
}

func (bd *DBWrapper) WithUrlShortener() outbound.DbOpsFunc {
	return func(o *outbound.DbOps) {
		o.UrlShortenerDAO = NewUrlShorteningAdapter(bd)
	}
}

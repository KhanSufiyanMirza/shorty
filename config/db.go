package config

// TODO: move to config
type RDBConfigOptions struct {
	Host      string // e.g: hostname:port  // google.com:8008
	Username  string
	Password  string
	DBName    string
	AppName   string
	Models    []interface{}
	ErrorCode ErrorCode
}

// TODO: move these constant into constant pkg
type Error int

const (
	NewUniquenessConstraintViolationError Error = iota
	DatabaseConnectionTimeoutError
	DatabaseSessionKilledError
)

type ErrorCode interface {
	GetError(code string) Error
}

func newRDBConfigOptions(host, username, password, dbname, appname string, models ...interface{}) RDBConfigOptions {
	return RDBConfigOptions{
		Host:     host,
		DBName:   dbname,
		Username: username,
		Password: password,
		AppName:  appname,
		Models:   models,
	}
}

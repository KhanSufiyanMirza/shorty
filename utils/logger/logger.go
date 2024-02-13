package logger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm/logger"
)

type Logger interface {
	Info(module string, msg string, err error, fields ...any)
	Warn(module string, msg string, err error, fields ...any)
	Error(module string, msg string, err error, fields ...any)
	Debug(module string, msg string, err error, fields ...any)
	Fatal(module string, msg string, err error, fields ...any)
	Fields(fields ...any) *LoggingService
	GetDBLogger() *dblogger
	Opentelemetry() (*tracesdk.TracerProvider, error)
}

type LoggingService struct {
	level     LogLevel
	component string
	ctx       context.Context
	dblogger  *dblogger
	// fields - only map[string]interface{} and []interface{} are accepted. []interface{} must
	// alternate string keys and arbitrary values, and extraneous ones are ignored.
	fields interface{}
}

var _ Logger = new(LoggingService)

// NewLogger creates a new Logger instance.
// It takes 'level' as the logging level and 'component' as component/module.
// It returns a Logger interface and an error, if any.
func NewLogger(level, component string) (Logger, error) {
	ls := LoggingService{
		level:     LogLevel(level),
		component: component,
	}
	return ls, nil
}

type LogLevel string

const (
	InfoLevel  LogLevel = "info"
	DebugLevel LogLevel = "debug"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
	PanicLevel LogLevel = "panic"
	NoLevel    LogLevel = "nolevel"
	Disabled   LogLevel = "disabled"
	TraceLevel LogLevel = "trace"
)

var CallerSkipFrameCount = zerolog.CallerSkipFrameCount + 1

// Init initializes a new Logger using the provided parameters.
// Creating a LoggingService instance with the log level, application name, context, and database logger.
// Using Zerlog package, setting up the time format and marshalling the errors.
// Defining a CallerMarshalFunc to modify the file path for logging and returns the modified file path and line number.
// Setting the logging level based on the provided LogLevel.
// Finally, It returns a created logging service instance as a Logger interface.
func Init(ctx context.Context, level LogLevel, appname string) Logger {
	if level == DebugLevel {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	}

	log := &LoggingService{
		level:     level,
		component: appname,
		ctx:       ctx,
		dblogger:  &dblogger{},
	}

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z07:00"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		userHome, _ := os.UserHomeDir()
		file = strings.ReplaceAll(short, userHome, "~")
		return file + ":" + strconv.Itoa(line)
	}

	switch level {
	case PanicLevel:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case FatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case WarnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case TraceLevel:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case NoLevel:
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case Disabled:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return log
}

// GetDBLogger retrieves the dblogger instance.
func (ls LoggingService) GetDBLogger() *dblogger {
	return ls.dblogger
}

// Fields is a helper function to use a map or slice to set fields using type assertion.
// Only map[string]interface{} and []interface{} are accepted. []interface{} must
// alternate string keys and arbitrary values, and extraneous ones are ignored.
func (ls LoggingService) Fields(args ...any) *LoggingService {
	ls.fields = args
	return &ls
}

// Fatal logs a fatal-level message using the LoggingService.
// Firstly, Creating a new logger context with specific fields for the log entry.
// And Creating a new fatal-level log entry with stack trace, error, and message.
func (ls LoggingService) Fatal(module string, msg string, err error, fields ...any) {
	ctx := log.With().
		Str("component", ls.component).
		Str("module", module).
		Fields(ls.fields).
		Fields(fields).
		CallerWithSkipFrameCount(CallerSkipFrameCount).
		Logger().
		WithContext(ls.ctx)
	log.Ctx(ctx).Fatal().Err(err).Msg(msg)
}

// Error logs a Error-level message using the LoggingService.
// Firstly, Creating a new logger context with specific fields for the log entry.
// And Creating a new Error-level log entry with stack trace, error, and message.
func (ls LoggingService) Error(module string, msg string, err error, fields ...any) {
	ctx := log.With().
		Str("component", ls.component).
		Str("module", module).
		Fields(ls.fields).
		Fields(fields).
		CallerWithSkipFrameCount(CallerSkipFrameCount).
		Logger().
		WithContext(ls.ctx)
	log.Ctx(ctx).Info().Err(err).Msg(msg)
}

// Warn logs a warning-level message using the LoggingService.
// Firstly, Creating a new logger context with specific fields for the log entry.
// And then creating a new warning-level log entry with stack trace, error, and message.
func (ls LoggingService) Warn(module string, msg string, err error, fields ...any) {
	ctx := log.With().
		Str("component", ls.component).
		Str("module", module).
		Fields(ls.fields).
		Fields(fields).
		CallerWithSkipFrameCount(CallerSkipFrameCount).
		Logger().
		WithContext(ls.ctx)
	log.Ctx(ctx).Warn().Err(err).Msg(msg)
}

// Info logs a Info-level message using the LoggingService.
// Firstly, Creating a new logger context with specific fields for the log entry.
// And Creating a new Info-level log entry with stack trace, error, and message.
func (ls LoggingService) Info(module string, msg string, err error, fields ...any) {
	ctx := log.With().
		Str("component", ls.component).
		Str("module", module).
		Fields(ls.fields).Fields(fields).
		CallerWithSkipFrameCount(CallerSkipFrameCount).
		Logger().
		WithContext(ls.ctx)
	log.Ctx(ctx).Info().Err(err).Msg(msg)
}

// Debug logs a debug-level message using the LoggingService.
// Firstly, Creating a new logger context with specific fields for the log entry.
// And then Create a new debug-level log entry with stack trace, error, and message.
func (ls LoggingService) Debug(module string, msg string, err error, fields ...any) {
	ctx := log.With().
		Str("component", ls.component).
		Str("module", module).
		Fields(ls.fields).
		Fields(fields).
		CallerWithSkipFrameCount(CallerSkipFrameCount).
		Logger().
		WithContext(ls.ctx)
	log.Ctx(ctx).Debug().Err(err).Msg(msg)
}

type dblogger struct{}

// GetDBLogger retrieves the dblogger instance.
func GetDBLogger() *dblogger {
	return &dblogger{}
}

// LogMode sets the log mode of the dblogger.
// It takes a logger.LogLevel parameter and returns a logger.Interface.
func (dbl dblogger) LogMode(logger.LogLevel) logger.Interface {
	return dbl
}

// Error method is used for logging error messages using the dblogger.
func (dbl dblogger) Error(ctx context.Context, msg string, opts ...interface{}) {
	log.Ctx(ctx).Error().Msg(fmt.Sprintf(msg, opts...))
}

// Warn method used to logs a warning message using the dblogger.
func (dbl dblogger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	log.Ctx(ctx).Warn().Msg(fmt.Sprintf(msg, opts...))
}

// Info method used to logs an informational message using the dblogger.
func (dbl dblogger) Info(ctx context.Context, msg string, opts ...interface{}) {
	log.Ctx(ctx).Info().Msg(fmt.Sprintf(msg, opts...))
}

// Trace function, logs a trace-level message using the dblogger.
// Firstly, Retrieve's the zerolog logger instance.
// And then, Checking if there's an error; set log level accordingly (Debug for errors, Trace).
// Assigning the key for duration based on the zerolog Duration Field Unit.
// It Log's an error if an unknown DurationFieldUnit value is noted.
// Calculate the duration and add it to the log event.
// Retrieve SQL query and rows count.
// Log's the SQL query if present and also, Log's the rows count if present and valid.
func (dbl dblogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	zl := log.Ctx(ctx)
	var event *zerolog.Event

	if err != nil {
		event = zl.Debug()
	} else {
		event = zl.Trace()
	}

	var dur_key string

	switch zerolog.DurationFieldUnit {
	case time.Nanosecond:
		dur_key = "elapsed_ns"
	case time.Microsecond:
		dur_key = "elapsed_us"
	case time.Millisecond:
		dur_key = "elapsed_ms"
	case time.Second:
		dur_key = "elapsed"
	case time.Minute:
		dur_key = "elapsed_min"
	case time.Hour:
		dur_key = "elapsed_hr"
	default:
		zl.Error().Interface("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).Msg("gormzerolog encountered a mysterious, unknown value for DurationFieldUnit")
		dur_key = "elapsed_"
	}

	event.Dur(dur_key, time.Since(begin))

	sql, rows := f()
	if sql != "" {
		event.Str("sql", sql)
	}
	if rows > -1 {
		event.Int64("rows", rows)
	}

	event.Send()

	return
}

func (ls LoggingService) Opentelemetry() (*tracesdk.TracerProvider, error) {
	return nil, nil
}

package logger

import (
	"fmt"
)

const (
	LogLevelKeyName = "LOG_LEVEL"
)

var _ fmt.Stringer = zapEncodingType("")

type zapEncodingType string

func (z zapEncodingType) String() string {
	switch z {
	case zapConsoleEncoding, zapJsonEncoding:
		return string(z)
	default:
		return string(defaultZapEncoding)
	}
}

const (
	zapJsonEncoding    zapEncodingType = "json"
	zapConsoleEncoding                 = "console"
)
const zapLogImplementor = "zap"
const bufferedLogImplementor = "buffered"
const noopLogImplementor = "noop"

const defaultLogLevel = string(DebugLevelString)
const defaultLogImplementer = zapLogImplementor
const defaultIsColored = true
const defaultZapEncoding = zapJsonEncoding

type loggerConfig struct {
	loggerLevelString string
	logImplementer    string
	//goEnv             string
	isCallerOn  bool
	zapEncoding zapEncodingType
	isColored   bool // isColored depends on the implementor
	printToIO   bool
}

type Option interface {
	apply(*loggerConfig)
}

type optionFunc func(*loggerConfig)

func (f optionFunc) apply(cfg *loggerConfig) {
	f(cfg)
}

func WithLoggerLevel(loggerLevel string) Option {
	return optionFunc(func(cfg *loggerConfig) {
		if _, err := parseLevelString(loggerLevel); err == nil {
			cfg.loggerLevelString = loggerLevel
			return
		}
		cfg.loggerLevelString = defaultLogLevel
	})
}

func WithZapColored() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.isColored = true
	})
}

func WithZapLoggerImplementer() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.logImplementer = zapLogImplementor
	})
}

func WithBufferedLoggerImplementer() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.logImplementer = bufferedLogImplementor
	})
}

func WithNoopImplementer() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.logImplementer = noopLogImplementor
	})
}

func WithPrintToIO() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.printToIO = true
	})
}

func WithZapConsoleEncoding() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.zapEncoding = zapConsoleEncoding
	})
}

func WithZapJsonEncoding() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.zapEncoding = zapJsonEncoding
	})
}

func WithCallerOn() Option {
	return optionFunc(func(cfg *loggerConfig) {
		cfg.isCallerOn = true
	})
}

package mysql

import (
	"os"
	"strconv"
	"time"
)

const (
	DefaultConnectionTimeout = 5
	DefaultMaxIdleTime       = 5
	DefaultTickerDuration    = 15
	DefaultTickerTime        = 15
	DefaultMaxConnections    = 10
	DefaultIdleConnections   = 10
)

type Config struct {
	Host                     string
	Port                     string
	DbName                   string
	Username                 string
	Password                 string
	MaxConnectionTimeout     time.Duration
	MaxIdleTime              time.Duration
	MaxConnections           int
	MaxIdleConnections       int
	ConnectionTickerDuration time.Duration
	ConnectionTickerTime     time.Duration
}

func LoadConfig() *Config {
	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	config.MaxConnectionTimeout = time.Minute *
		time.Duration(getValueFromOsEnv("DB_TIMEOUT_MINUTES", DefaultConnectionTimeout))

	config.MaxIdleTime = time.Minute *
		time.Duration(getValueFromOsEnv("DB_IDLE_TIME_MINUTES", DefaultMaxIdleTime))

	config.ConnectionTickerDuration = time.Second *
		time.Duration(getValueFromOsEnv("DB_TICKER_DURATION_SECONDS", DefaultTickerDuration))

	config.ConnectionTickerTime = time.Minute *
		time.Duration(getValueFromOsEnv("DB_TICKER_TIME_MINUTES", DefaultTickerTime))

	config.MaxConnections = DefaultMaxConnections
	config.MaxIdleConnections = DefaultIdleConnections

	return config
}

func getValueFromOsEnv(name string, defaultValue int) int {
	if os.Getenv(name) != "" {
		value, err := strconv.Atoi(os.Getenv(name))

		if err != nil {
			return defaultValue
		}

		return value
	} else {
		return defaultValue
	}
}

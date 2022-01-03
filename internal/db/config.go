package db

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host string
	Port string
	DbName string
	Username string
	Password string
	MaxConnectionTimeout time.Duration
	MaxIdleTime time.Duration
	MaxConnections int
	MaxIdleConnections int
}

func LoadConfig() *Config {
	config := &Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if os.Getenv("DB_TIMEOUT_MINUTES") != "" {
		minutes, err := strconv.Atoi(os.Getenv("DB_TIMEOUT_MINUTES"))

		if err != nil {
			config.MaxConnectionTimeout = time.Minute * 3
		}

		config.MaxConnectionTimeout = time.Minute * time.Duration(minutes)
	} else {
		config.MaxConnectionTimeout = time.Minute * 3
	}

	if os.Getenv("DB_IDLE_TIME_SECONDS") != "" {
		seconds, err := strconv.Atoi(os.Getenv("DB_IDLE_TIME_SECONDS"))

		if err != nil {
			config.MaxIdleTime = time.Second * 10
		}

		config.MaxIdleTime = time.Second * time.Duration(seconds)
	} else {
		config.MaxIdleTime = time.Second * 10
	}

	config.MaxConnections = 10
	config.MaxIdleConnections = 10

	return config
}

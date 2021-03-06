package app

import (
	"fmt"
	"github.com/aerostatka/db-testing/internal/action"
	"github.com/aerostatka/db-testing/internal/mysql"
	service2 "github.com/aerostatka/db-testing/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start(mode string) {
	logger := getLogger()

	dbConfig := mysql.LoadConfig()
	dbClient, err := getDbClient(dbConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbClient.Close()

	dbRep := mysql.CreateDbConnectionRepository(dbClient, logger, dbConfig)
	service := service2.CreateDefaultConnectionService(dbRep, logger)

	factory := action.CreateDefaultFactory(service, logger)
	action, err := factory.GetAction(mode)

	if err != nil {
		logger.Fatal(err.Error())
	}

	err = action.SanityCheck()

	if err != nil {
		logger.Fatal(err.Error())
	}

	result := action.Perform()

	if result.Result {
		logger.Info(
			fmt.Sprintf("Action %s is executed successfully", mode),
		)
	} else {
		logger.Info(
			fmt.Sprintf("Action %s is executed with an error", mode),
		)
	}
}

func getLogger() *zap.Logger {
	var err error

	config := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	return log
}

func getDbClient(dbConfig *mysql.Config) (*sqlx.DB, error) {
	client, err := sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DbName,
		),
	)

	if err != nil {
		if client != nil {
			defer client.Close()
		}
		return nil, err
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(dbConfig.MaxConnectionTimeout)
	client.SetMaxOpenConns(dbConfig.MaxConnections)
	client.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	client.SetConnMaxIdleTime(dbConfig.MaxIdleTime)

	return client, nil
}

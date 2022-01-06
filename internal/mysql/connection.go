package mysql

import (
	"fmt"
	"github.com/aerostatka/db-testing/internal/domains"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type ConnectionRepository struct {
	client *sqlx.DB
	logger *zap.Logger
	config *Config
}

func CreateDbConnectionRepository(c *sqlx.DB, log *zap.Logger, conf *Config) *ConnectionRepository {
	return &ConnectionRepository{
		client: c,
		logger: log,
		config: conf,
	}
}

func (rep *ConnectionRepository) TestConnection() bool {
	ticker := time.NewTicker(rep.config.ConnectionTickerDuration)
	done := make(chan bool)

	go rep.showProcesses(done, ticker)

	time.Sleep(rep.config.ConnectionTickerTime)
	ticker.Stop()
	done <- true
	rep.logger.Info("Ticker stopped")

	return true
}

func (rep *ConnectionRepository) showProcesses(done chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			rep.logger.Info("-----------------")
			rep.logger.Info(fmt.Sprintf("Tick at %v", t))

			var processes []domains.Process
			err := rep.client.Select(&processes, "SELECT * FROM INFORMATION_SCHEMA.PROCESSLIST")

			if err != nil {
				rep.logger.Debug(fmt.Sprintf("Error during mysql connection: %s", err.Error()))
			}

			rep.logger.Info("Processes list")
			for _, process := range processes {
				rep.logger.Info(
					fmt.Sprintf(
						"ID %d: USER %s, DB %s, COMMAND %s, TIME %d, INFO '%s'",
						process.Id,
						process.User.String,
						process.Db.String,
						process.Command.String,
						process.Time,
						process.Info.String,
					),
				)
			}
		}
	}
}

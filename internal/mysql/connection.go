package mysql

import (
	"fmt"
	"github.com/aerostatka/db-testing/internal/domains"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ConnectionRepository struct {
	client *sqlx.DB
	logger *zap.Logger
}

func CreateDbConnectionRepository(c *sqlx.DB, log *zap.Logger) *ConnectionRepository {
	return &ConnectionRepository{
		client: c,
		logger: log,
	}
}

func (rep *ConnectionRepository) TestConnection() bool {
	processes := []domains.Process{}
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

	return true
}

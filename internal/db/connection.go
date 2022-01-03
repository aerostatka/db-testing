package db

import (
	"fmt"
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
	_, err := rep.client.Exec("SHOW processlist")

	if err != nil {
		rep.logger.Debug(fmt.Sprintf("Error during db connection: %s", err.Error()))
		return false
	}

	return true
}

package service

import (
	"github.com/aerostatka/db-testing/internal/mysql"
	"go.uber.org/zap"
)

type ConnectionRepository interface {
	TestConnection() bool
}

type ConnectionTestService interface {
	TestDbConnection() bool
	TestDynamoDbConnection() bool
}

type DefaultConnectionService struct {
	logger *zap.Logger
	dbRep  ConnectionRepository
}

func CreateDefaultConnectionService(rep *mysql.ConnectionRepository, log *zap.Logger) *DefaultConnectionService {
	return &DefaultConnectionService{
		logger: log,
		dbRep:  rep,
	}
}

func (service *DefaultConnectionService) TestDbConnection() bool {
	return service.dbRep.TestConnection()
}

func (service *DefaultConnectionService) TestDynamoDbConnection() bool {
	return true
}

package action

import (
	"errors"
	"fmt"
	"github.com/aerostatka/db-testing/internal/domains"
	"go.uber.org/zap"
)

const (
	ConsoleModeDbTest       = "db-test"
	ConsoleModeDynamoDbTest = "dynamo-test"
)

type ConsoleAction interface {
	SanityCheck() error
	Perform() *domains.ConnectivityResult
}

type FactoryInterface interface {
	GetAction(mode string) (ConsoleAction, error)
}

type DefaultFactory struct {
	service domains.ConnectionTestService
	log     *zap.Logger
}

func (factory *DefaultFactory) GetAction(mode string) (ConsoleAction, error) {
	switch mode {
	case ConsoleModeDbTest:
		return createDbAction(factory.service, factory.log), nil
	default:
		return nil, errors.New(fmt.Sprintf("Action %s is not found", mode))
	}
}

func CreateDefaultFactory(s domains.ConnectionTestService, logger *zap.Logger) *DefaultFactory {
	return &DefaultFactory{
		service: s,
		log:     logger,
	}
}

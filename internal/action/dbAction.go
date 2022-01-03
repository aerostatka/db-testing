package action

import (
	"errors"
	"github.com/aerostatka/db-testing/internal/domains"
	"github.com/aerostatka/db-testing/internal/service"
	"go.uber.org/zap"
	"os"
)

type DbAction struct {
	service service.ConnectionTestService
	log     *zap.Logger
}

func createDbAction(s service.ConnectionTestService, logger *zap.Logger) *DbAction {
	return &DbAction{
		service: s,
		log:     logger,
	}
}

func (action *DbAction) SanityCheck() error {
	if os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_USERNAME") == "" {
		return errors.New("Required mysql parameters are not specified")
	}

	return nil
}

func (action *DbAction) Perform() *domains.ConnectivityResult {
	return &domains.ConnectivityResult{
		Result: action.service.TestDbConnection(),
	}
}

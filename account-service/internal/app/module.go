package app

import (
	"github.com/saufiroja/sosmed-app/account-service/config"
	"github.com/saufiroja/sosmed-app/account-service/internal/controllers"
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers"
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers/command"
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers/query"
	"github.com/saufiroja/sosmed-app/account-service/internal/repositories"
	"github.com/saufiroja/sosmed-app/account-service/internal/services"
	"github.com/saufiroja/sosmed-app/account-service/pkg/database"
	"github.com/sirupsen/logrus"
)

type Module struct {
	controllers.AccountControllerInterface
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(conf *config.AppConfig, logger *logrus.Logger) {
	// register database
	db := database.NewPostgres(conf, logger)
	// register repository
	repo := repositories.NewAccountRepository(db)
	// register service
	service := services.NewAccountService(repo, db, logger)
	// register command
	insertUserCommand := command.NewInsertUserCommandHandler(service)
	command := command.NewCommand(insertUserCommand)
	// register query
	getAccountByUsernameAndEmailQuery := query.NewGetAccountByUsernameAndEmailHandlerQuery(service, logger)
	query := query.NewQuery(getAccountByUsernameAndEmailQuery)

	// register handler
	handler := handlers.NewHandler(command, query)

	// register controller
	controller := controllers.NewAccountController(handler)

	m.AccountControllerInterface = controller
}

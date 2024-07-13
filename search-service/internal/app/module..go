package app

import (
	"github.com/saufiroja/sosmed/search-service/config"
	consumer "github.com/saufiroja/sosmed/search-service/internal/consumers"
	"github.com/saufiroja/sosmed/search-service/internal/controllers"
	"github.com/saufiroja/sosmed/search-service/internal/handlers"
	"github.com/saufiroja/sosmed/search-service/internal/handlers/query"
	"github.com/saufiroja/sosmed/search-service/internal/repositories"
	"github.com/saufiroja/sosmed/search-service/internal/services"
	"github.com/saufiroja/sosmed/search-service/pkg/database"
	"github.com/saufiroja/sosmed/search-service/pkg/messaging"
	"github.com/sirupsen/logrus"
)

type Module struct {
	controllers.SearchControllerInterface
	*consumer.Consumer
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(conf *config.AppConfig, logger *logrus.Logger) {
	// Register the database
	db := database.NewElasticsearch(conf, logger)
	// Register the repository
	userRepository := repositories.NewUserRepository(db)
	// Register the service
	userService := services.NewUserService(userRepository, logger)

	// message broker
	kafkaConsumer := messaging.NewKafkaConsumer(conf, logger)
	consumer := consumer.NewConsumer(kafkaConsumer, userService, logger)

	// query handler
	getUserByUsernameQuery := query.NewGetUserByUsernameQueryHandler(userService)
	getAllUsersQuery := query.NewGetAllUsersQueryHandler(userService)
	query := query.NewQuery(getUserByUsernameQuery, getAllUsersQuery)

	// Register the handler
	handler := handlers.NewHandler(query)

	// Register the controller
	searchController := controllers.NewController(handler)

	m.SearchControllerInterface = searchController
	m.Consumer = consumer
}

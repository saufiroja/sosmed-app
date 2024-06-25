package app

import (
	"github.com/saufiroja/sosmed-app/auth-service/config"
	"github.com/saufiroja/sosmed-app/auth-service/internal/controllers"
	internalGrpc "github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/handlers"
	"github.com/saufiroja/sosmed-app/auth-service/internal/handlers/command"
	"github.com/saufiroja/sosmed-app/auth-service/internal/repositories"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
	"github.com/saufiroja/sosmed-app/auth-service/internal/utils"
	"github.com/saufiroja/sosmed-app/auth-service/pkg/database"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

type Module struct {
	controllers.AuthControllerInterface
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(conf *config.AppConfig, logger *logrus.Logger, conn *grpc.ClientConn) {
	// google auth
	googleAuth := utils.NewGoogleAuth(conf)
	// utils
	generateToken := utils.NewGenerateToken(conf)
	// client
	accountClient := internalGrpc.NewAccountServiceClient(conn)
	// Register the database
	db := database.NewPostgres(conf, logger)
	// Register the repository
	authRepository := repositories.NewAuthRepository(db)
	// Register the service
	authService := services.NewAuthService(authRepository, logger, db, generateToken, accountClient, googleAuth, conf)
	// handler command
	registerHandler := command.NewRegisterCommandHandler(authService, logger)
	loginHandler := command.NewLoginCommandHandler(authService, logger)
	refreshTokenHandler := command.NewRefreshTokenCommandHandler(authService, logger)
	googleAuthHandler := command.NewGoogleAuthCommandHandler(authService)
	googleAuthCallbackHandler := command.NewGoogleAuthCallbackCommandHandler(authService)

	// Register the command
	command := command.NewCommand(
		registerHandler,
		loginHandler,
		refreshTokenHandler,
		googleAuthHandler,
		googleAuthCallbackHandler,
	)
	// handler
	handler := handlers.NewHandler(command)
	// Register the controller
	controller := controllers.NewController(handler)

	m.AuthControllerInterface = controller
}

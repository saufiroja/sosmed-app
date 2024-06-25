package command

import (
	"context"
	"errors"

	"github.com/saufiroja/sosmed-app/auth-service/internal/contracts/requests"
	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
	"github.com/sirupsen/logrus"
)

type LoginCommandHandler struct {
	authService services.AuthServiceInterface
	logger      *logrus.Logger
}

func NewLoginCommandHandler(authService services.AuthServiceInterface, logger *logrus.Logger) *LoginCommandHandler {
	return &LoginCommandHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, command *grpc.LoginRequest) (*grpc.Response, error) {
	request := &requests.LoginRequest{
		Email:       command.Email,
		Username:    command.Username,
		Password:    command.Password,
		AccountType: command.AccountType,
	}

	user, err := h.authService.Login(ctx, request)
	if err != nil {
		h.logger.Error(err)
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

package command

import (
	"context"
	"errors"

	"github.com/saufiroja/sosmed-app/auth-service/internal/contracts/requests"
	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
	"github.com/sirupsen/logrus"
)

type RegisterCommandHandler struct {
	authService services.AuthServiceInterface
	logger      *logrus.Logger
}

func NewRegisterCommandHandler(authService services.AuthServiceInterface, logger *logrus.Logger) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *RegisterCommandHandler) Handle(ctx context.Context, command *grpc.RegisterRequest) (*grpc.Empty, error) {
	accountType, err := h.authService.GetAccountType(ctx, command.AccountType)
	if err != nil {
		h.logger.Error(err)
		return nil, errors.New("account type not found")
	}

	request := &requests.RegisterRequest{
		Username:    command.Username,
		Email:       command.Email,
		Password:    command.Password,
		FullName:    command.FullName,
		AccountType: accountType.ID,
	}

	err = h.authService.Register(ctx, request)
	if err != nil {
		h.logger.Error(err)
		return nil, errors.New("failed to register user")
	}

	return &grpc.Empty{}, nil
}

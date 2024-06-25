package command

import (
	"context"

	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
	"github.com/sirupsen/logrus"
)

type RefreshTokenCommandHandler struct {
	authService services.AuthServiceInterface
	logger      *logrus.Logger
}

func NewRefreshTokenCommandHandler(authService services.AuthServiceInterface, logger *logrus.Logger) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *RefreshTokenCommandHandler) Handle(ctx context.Context, command *grpc.RefreshTokenRequest) (*grpc.Response, error) {
	result, err := h.authService.RefreshToken(ctx, command.RefreshToken)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	return result, nil
}

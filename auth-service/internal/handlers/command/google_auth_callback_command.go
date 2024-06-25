package command

import (
	"context"

	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
)

type GoogleAuthCallbackCommandHandler struct {
	authService services.AuthServiceInterface
}

func NewGoogleAuthCallbackCommandHandler(authService services.AuthServiceInterface) *GoogleAuthCallbackCommandHandler {
	return &GoogleAuthCallbackCommandHandler{
		authService: authService,
	}
}

func (c *GoogleAuthCallbackCommandHandler) Handle(ctx context.Context, request *grpc.GoogleAuthCallbackRequest) (*grpc.Response, error) {
	res, err := c.authService.GoogleAuthCallback(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	return res, nil
}

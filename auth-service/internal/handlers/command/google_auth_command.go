package command

import (
	"context"

	internalGrpc "github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GoogleAuthCommandHandler struct {
	authService services.AuthServiceInterface
}

func NewGoogleAuthCommandHandler(authService services.AuthServiceInterface) *GoogleAuthCommandHandler {
	return &GoogleAuthCommandHandler{
		authService: authService,
	}
}

func (c *GoogleAuthCommandHandler) Handle(ctx context.Context, request *internalGrpc.Empty) (*internalGrpc.GoogleAuthResponse, error) {
	url, err := c.authService.GoogleAuth(ctx)
	if err != nil {
		return nil, err
	}

	header := metadata.Pairs("Location", *url)

	_ = grpc.SendHeader(ctx, header)

	return &internalGrpc.GoogleAuthResponse{
		Url: *url,
	}, nil
}

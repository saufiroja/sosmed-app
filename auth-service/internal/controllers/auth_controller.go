package controllers

import (
	"context"

	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/handlers"
)

type Controller struct {
	handlers *handlers.Handler
}

func NewController(handlers *handlers.Handler) AuthControllerInterface {
	return &Controller{
		handlers: handlers,
	}
}

func (c *Controller) Login(ctx context.Context, request *grpc.LoginRequest) (*grpc.Response, error) {
	return c.handlers.Command.LoginCommandHandler.Handle(ctx, request)
}

func (c *Controller) RefreshToken(ctx context.Context, request *grpc.RefreshTokenRequest) (*grpc.Response, error) {
	return c.handlers.Command.RefreshTokenCommandHandler.Handle(ctx, request)
}

func (c *Controller) Register(ctx context.Context, request *grpc.RegisterRequest) (*grpc.Empty, error) {
	return c.handlers.Command.RegisterCommandHandler.Handle(ctx, request)
}

func (c *Controller) GoogleAuth(ctx context.Context, request *grpc.Empty) (*grpc.GoogleAuthResponse, error) {
	return c.handlers.Command.GoogleAuthCommandHandler.Handle(ctx, request)
}

func (c *Controller) GoogleAuthCallback(ctx context.Context, request *grpc.GoogleAuthCallbackRequest) (*grpc.Response, error) {
	return c.handlers.Command.GoogleAuthCallbackCommandHandler.Handle(ctx, request)
}

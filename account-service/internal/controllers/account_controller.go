package controllers

import (
	"context"

	"github.com/saufiroja/sosmed-app/account-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/account-service/internal/handlers"
)

type AccountController struct {
	handler *handlers.Handler
}

func NewAccountController(handler *handlers.Handler) AccountControllerInterface {
	return &AccountController{
		handler: handler,
	}
}

func (a *AccountController) GetAccountByEmailAndUsername(ctx context.Context, request *grpc.GetAccountByEmailAndUsernameRequest) (*grpc.GetAccountByEmailAndUsernameResponse, error) {
	return a.handler.Query.GetAccountByUsernameAndEmailHandlerQuery.Handle(ctx, request.Username, request.Email)
}

func (a *AccountController) InsertUser(ctx context.Context, request *grpc.InsertUserRequest) (*grpc.Empty, error) {
	return a.handler.Command.InsertUserCommandHandler.Handle(ctx, request)
}

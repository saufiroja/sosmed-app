package controllers

import (
	"context"

	"github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/saufiroja/sosmed/search-service/internal/handlers"
)

type Controller struct {
	handler *handlers.Handler
}

func NewController(handler *handlers.Handler) SearchControllerInterface {
	return &Controller{handler: handler}
}

func (c *Controller) SearchUserByUsername(ctx context.Context, req *grpc.SearchUserByUsernameRequest) (*grpc.SearchUserByUsernameResponse, error) {
	return c.handler.Query.GetUserByUsernameQueryHandler.Handle(ctx, req)
}

func (c *Controller) SearchAllUsers(ctx context.Context, req *grpc.SearchAllUsersRequest) (*grpc.SearchAllUsersResponse, error) {
	return c.handler.Query.GetAllUsersQueryHandler.Handle(ctx, req)
}

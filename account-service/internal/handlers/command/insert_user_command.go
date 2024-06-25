package command

import (
	"context"

	"github.com/saufiroja/sosmed-app/account-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/account-service/internal/models"
	"github.com/saufiroja/sosmed-app/account-service/internal/services"
)

type InsertUserCommandHandler struct {
	accountService services.AccountServiceInterface
}

func NewInsertUserCommandHandler(accountService services.AccountServiceInterface) *InsertUserCommandHandler {
	return &InsertUserCommandHandler{
		accountService: accountService,
	}
}

func (i *InsertUserCommandHandler) Handle(ctx context.Context, request *grpc.InsertUserRequest) (*grpc.Empty, error) {
	user := &models.User{
		UserID:   request.UserId,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		FullName: request.FullName,
	}

	err := i.accountService.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}

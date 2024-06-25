package clients

import (
	"context"

	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	internalGrpc "github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
)

type AccountClient struct {
	accountClient internalGrpc.AccountServiceClient
}

func NewAccountClient(accountClient internalGrpc.AccountServiceClient) *AccountClient {
	return &AccountClient{
		accountClient: accountClient,
	}
}

func (a *AccountClient) InsertUser(ctx context.Context, request *grpc.InsertUserRequest) (*grpc.Empty, error) {
	return a.accountClient.InsertUser(ctx, request)
}

func (a *AccountClient) GetAccountByEmailAndUsername(ctx context.Context, request *grpc.GetAccountByEmailAndUsernameRequest) (*grpc.GetAccountByEmailAndUsernameResponse, error) {
	return a.accountClient.GetAccountByEmailAndUsername(ctx, request)
}

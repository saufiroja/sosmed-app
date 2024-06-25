package services

import (
	"context"

	"github.com/saufiroja/sosmed-app/auth-service/internal/contracts/requests"
	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/models"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, request *requests.RegisterRequest) error
	GetAccountType(ctx context.Context, accountType string) (*models.AccountType, error)
	Login(ctx context.Context, request *requests.LoginRequest) (*grpc.Response, error)
	RefreshToken(ctx context.Context, refreshToken string) (*grpc.Response, error)
	GoogleAuth(ctx context.Context) (*string, error)
	GoogleAuthCallback(ctx context.Context, code string) (*grpc.Response, error)
}

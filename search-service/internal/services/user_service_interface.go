package services

import (
	"context"

	"github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/saufiroja/sosmed/search-service/internal/models"
)

type UserServiceInterface interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAllUsers(ctx context.Context, request *grpc.SearchAllUsersRequest) ([]models.User, error)
}

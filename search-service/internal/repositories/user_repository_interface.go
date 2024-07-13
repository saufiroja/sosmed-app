package repositories

import (
	"context"

	"github.com/saufiroja/sosmed/search-service/internal/models"
)

type UserRepositoryInterface interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAllUsers(ctx context.Context, page, limit *int32) ([]models.User, error)
}

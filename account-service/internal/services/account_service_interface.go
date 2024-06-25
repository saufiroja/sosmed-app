package services

import (
	"context"

	"github.com/saufiroja/sosmed-app/account-service/internal/models"
)

type AccountServiceInterface interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*models.Account, error)
}

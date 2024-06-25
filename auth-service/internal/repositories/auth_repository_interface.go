package repositories

import (
	"context"
	"database/sql"

	"github.com/saufiroja/sosmed-app/auth-service/internal/models"
)

type AuthRepositoryInterface interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *models.User) error
	GetAccountType(ctx context.Context, accountType string) (*models.AccountType, error)
}

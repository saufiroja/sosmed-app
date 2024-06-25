package repositories

import (
	"context"
	"database/sql"

	"github.com/saufiroja/sosmed-app/account-service/internal/models"
)

type AccountRepositoryInterface interface {
	InsertAccount(ctx context.Context, tx *sql.Tx, account *models.Account) error
	InsertProfile(ctx context.Context, tx *sql.Tx, profile *models.Profile) error
	GetAccountByUsernameAndEmail(ctx context.Context, username, email string) (*models.Account, error)
}

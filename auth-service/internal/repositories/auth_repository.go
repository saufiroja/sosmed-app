package repositories

import (
	"context"
	"database/sql"

	"github.com/saufiroja/sosmed-app/auth-service/internal/models"
	"github.com/saufiroja/sosmed-app/auth-service/pkg/database"
)

type authRepository struct {
	Db database.PostgresInterface
}

func NewAuthRepository(db database.PostgresInterface) AuthRepositoryInterface {
	return &authRepository{Db: db}
}

func (a *authRepository) InsertUser(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `INSERT INTO users 
	(user_id, account_type_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4)`

	_, err := tx.ExecContext(ctx,
		query,
		user.UserID,
		user.AccountTypeID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (a *authRepository) GetAccountType(ctx context.Context, accountType string) (*models.AccountType, error) {
	db := a.Db.DbConnection()

	query := `SELECT id, account_name FROM account_type WHERE account_name = $1`

	account := &models.AccountType{}
	err := db.QueryRowContext(ctx, query, accountType).Scan(&account.ID, &account.AccountName)
	if err != nil {
		return nil, err
	}

	return account, nil
}

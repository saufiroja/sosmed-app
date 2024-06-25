package repositories

import (
	"context"
	"database/sql"

	"github.com/saufiroja/sosmed-app/account-service/internal/models"
	"github.com/saufiroja/sosmed-app/account-service/pkg/database"
)

type accountRepository struct {
	Db database.PostgresInterface
}

func NewAccountRepository(db database.PostgresInterface) AccountRepositoryInterface {
	return &accountRepository{
		Db: db,
	}
}

func (r *accountRepository) InsertAccount(ctx context.Context, tx *sql.Tx, account *models.Account) error {
	query := `INSERT INTO accounts (account_id, user_id, full_name, username, email, password, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := tx.ExecContext(ctx, query, account.AccountID, account.UserID, account.FullName, account.Username, account.Email, account.Password, account.CreatedAt, account.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) InsertProfile(ctx context.Context, tx *sql.Tx, profile *models.Profile) error {
	query := `INSERT INTO profiles (profile_id, user_id, avatar, bio, location, website, birth_date, phone_number, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := tx.ExecContext(ctx, query, profile.ProfileID, profile.UserID, nil, nil, nil, nil, nil, nil, profile.CreatedAt, profile.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) GetAccountByUsernameAndEmail(ctx context.Context, username, email string) (*models.Account, error) {
	db := r.Db.DbConnection()

	query := `SELECT account_id, user_id, full_name, username, email, password FROM accounts WHERE username=$1 AND email=$2`

	row := db.QueryRowContext(ctx, query, username, email)

	var account models.Account
	err := row.Scan(&account.AccountID, &account.UserID, &account.FullName, &account.Username, &account.Email, &account.Password)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

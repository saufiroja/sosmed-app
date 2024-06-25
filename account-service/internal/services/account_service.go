package services

import (
	"context"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/saufiroja/sosmed-app/account-service/internal/models"
	"github.com/saufiroja/sosmed-app/account-service/internal/repositories"
	"github.com/saufiroja/sosmed-app/account-service/pkg/database"
	"github.com/sirupsen/logrus"
)

type accountService struct {
	accountRepo repositories.AccountRepositoryInterface
	db          database.PostgresInterface
	logger      *logrus.Logger
}

func NewAccountService(accountRepo repositories.AccountRepositoryInterface, db database.PostgresInterface, logger *logrus.Logger) AccountServiceInterface {
	return &accountService{
		accountRepo: accountRepo,
		db:          db,
		logger:      logger,
	}
}

func (s *accountService) InsertUser(ctx context.Context, user *models.User) error {
	s.logger.Info(fmt.Sprintf("inserting user with username: %s and email: %s", user.Username, user.Email))

	// start transaction
	tx := s.db.StartTransaction(ctx)

	account := &models.Account{
		AccountID: ulid.MustNew(ulid.Now(), nil).String(),
		UserID:    user.UserID,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// insert account
	err := s.accountRepo.InsertAccount(ctx, tx, account)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error inserting account: %v", err))
		s.db.RollbackTransaction(tx)
		return err
	}

	profile := &models.Profile{
		ProfileID: ulid.MustNew(ulid.Now(), nil).String(),
		UserID:    account.UserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// insert profile
	err = s.accountRepo.InsertProfile(ctx, tx, profile)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error inserting profile: %v", err))
		s.db.RollbackTransaction(tx)
		return err
	}

	// commit transaction
	s.db.CommitTransaction(tx)

	return nil
}

func (s *accountService) GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*models.Account, error) {
	s.logger.Info(fmt.Sprintf("getting user with username: %s and email: %s", username, email))
	account, err := s.accountRepo.GetAccountByUsernameAndEmail(ctx, username, email)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting account: %v", err))
		return nil, err
	}

	return account, nil
}

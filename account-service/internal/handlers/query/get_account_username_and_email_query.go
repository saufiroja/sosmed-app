package query

import (
	"context"

	"github.com/saufiroja/sosmed-app/account-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/account-service/internal/mappers"
	"github.com/saufiroja/sosmed-app/account-service/internal/services"
	"github.com/sirupsen/logrus"
)

type GetAccountByUsernameAndEmailHandlerQuery struct {
	accountService services.AccountServiceInterface
	logger         *logrus.Logger
}

func NewGetAccountByUsernameAndEmailHandlerQuery(accountService services.AccountServiceInterface, logger *logrus.Logger) *GetAccountByUsernameAndEmailHandlerQuery {
	return &GetAccountByUsernameAndEmailHandlerQuery{
		accountService: accountService,
		logger:         logger,
	}
}

func (g *GetAccountByUsernameAndEmailHandlerQuery) Handle(ctx context.Context, username, email string) (*grpc.GetAccountByEmailAndUsernameResponse, error) {
	account, err := g.accountService.GetUserByUsernameAndEmail(ctx, username, email)
	if err != nil {
		g.logger.Error(err)
		return nil, err
	}

	res := mappers.ToGetAccountByEmailAndUsernameResponse(account)

	return res, nil
}

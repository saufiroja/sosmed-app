package mappers

import (
	"github.com/saufiroja/sosmed-app/account-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/account-service/internal/models"
)

func ToGetAccountByEmailAndUsernameResponse(account *models.Account) *grpc.GetAccountByEmailAndUsernameResponse {
	return &grpc.GetAccountByEmailAndUsernameResponse{
		UserId:   account.UserID,
		Username: account.Username,
		Email:    account.Email,
		FullName: account.FullName,
		Password: account.Password,
	}
}

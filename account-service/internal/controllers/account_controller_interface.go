package controllers

import "github.com/saufiroja/sosmed-app/account-service/internal/grpc"

type AccountControllerInterface interface {
	grpc.AccountServiceServer
}

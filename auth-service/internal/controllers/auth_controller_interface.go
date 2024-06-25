package controllers

import "github.com/saufiroja/sosmed-app/auth-service/internal/grpc"

type AuthControllerInterface interface {
	grpc.AuthServiceServer
}

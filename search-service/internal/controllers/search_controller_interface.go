package controllers

import "github.com/saufiroja/sosmed/search-service/internal/grpc"

type SearchControllerInterface interface {
	grpc.SearchServiceServer
}

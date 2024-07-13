package app

import (
	"net"

	internalGrpc "github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	*grpc.Server
}

func NewGrpc(module *Module) *Grpc {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	internalGrpc.RegisterSearchServiceServer(grpcServer, module)

	return &Grpc{
		Server: grpcServer,
	}
}

func (g *Grpc) StartGrpc(lis *net.Listener, logger *logrus.Logger) {
	err := g.Serve(*lis)
	if err != nil {
		logger.Panic(err)
	}
}

func (g *Grpc) GrpcShutdown() {
	g.GracefulStop()
}

package app

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/saufiroja/sosmed-app/account-service/config"
	internalGrpc "github.com/saufiroja/sosmed-app/account-service/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Rest struct {
	*http.Server
}

func NewRest(listener net.Listener, ctx context.Context, grpcServer *Grpc, logger *logrus.Logger, conf *config.AppConfig) *Rest {
	conn, err := grpc.DialContext(
		ctx,
		listener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	gatewayMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher),
	)

	err = internalGrpc.RegisterAccountServiceHandler(ctx, gatewayMux, conn)
	if err != nil {
		logger.Panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", gatewayMux)

	restServer := &http.Server{
		Addr:    conf.Http.Port,
		Handler: mux,
	}

	return &Rest{
		Server: restServer,
	}
}

func (r *Rest) HttpStart(logger *logrus.Logger) {
	err := r.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return
	} else if err != nil {
		logger.Panic(err)
		return
	}
}

func (r *Rest) HttpStop(ctx context.Context, logger *logrus.Logger) {
	err := r.Shutdown(ctx)
	if err != nil {
		logger.Panic(err)
		return
	}
}

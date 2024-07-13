package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/saufiroja/sosmed/search-service/config"
	"github.com/sirupsen/logrus"
)

type App struct {
	rest   *Rest
	grpc   *Grpc
	logger *logrus.Logger
}

func NewApp() *App {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	return &App{
		logger: logger,
	}
}

func (a *App) RunApp() {
	conf := config.NewAppConfig()
	module := NewModule()

	httpPort := conf.Http.Port
	grpcPort := conf.Grpc.Port

	// initialize module
	module.Initialize(conf, a.logger)

	// start consumer
	module.Consumer.Start()

	// grpc
	grpcListen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		a.logger.Panic(err)
	}

	go func() {
		a.grpc = NewGrpc(module)
		a.grpc.StartGrpc(&grpcListen, a.logger)
	}()

	go func() {
		a.rest = NewRest(grpcListen, context.Background(), a.grpc, a.logger, conf)
		a.rest.HttpStart(a.logger)
	}()

	fmt.Println("---------------------------------------")
	fmt.Println("GRPC Server started on port:", grpcPort)
	fmt.Println("HTTP Server started on port:", httpPort)
	fmt.Println("---------------------------------------")

	os := map[string]operation{
		"grpc": func(ctx context.Context) error {
			a.grpc.GrpcShutdown()
			return nil
		},
		"http": func(ctx context.Context) error {
			return a.rest.Shutdown(ctx)
		},
	}

	// wait for shutdown signal
	<-a.Shutdown(context.Background(), 10*time.Second, os, a.logger)
}

type operation func(ctx context.Context) error

func (a *App) Shutdown(ctx context.Context, timeout time.Duration, ops map[string]operation, logger *logrus.Logger) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		// tunggu sinyal dari os
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

		// tunggu sinyal
		<-sign

		// set timeout untuk operasi yang belum selesai
		timeoutFunc := time.AfterFunc(timeout, func() {
			logger.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(1)
		})

		defer timeoutFunc.Stop()

		// tunggu semua operasi selesai
		var wg sync.WaitGroup

		// jalankan operasi
		for i, v := range ops {
			wg.Add(1)
			valueOs := v
			indexOs := i
			go func() {
				defer wg.Done()

				logger.Printf("cleaning up: %s", indexOs)
				if err := valueOs(ctx); err != nil {
					logger.Printf("%s: clean up failed: %v", indexOs, err)
					return
				}

				logger.Printf("%s was shutdown gracefully", indexOs)
			}()
		}

		wg.Wait()
		close(wait)

		logger.Println("shutdown completed")
	}()

	return wait
}

package config

import "os"

func initHttp(conf *AppConfig) {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	accountServiceHost := os.Getenv("ACCOUNT_SERVICE_URL")

	conf.Http.Port = httpPort
	conf.Grpc.Port = grpcPort
	conf.AccountService.URL = accountServiceHost
}

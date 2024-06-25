package config

import "os"

func initHttp(conf *AppConfig) {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "localhost:8082"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "localhost:50052"
	}

	conf.Http.Port = httpPort
	conf.Grpc.Port = grpcPort
}

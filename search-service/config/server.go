package config

import "os"

func initServer(conf *AppConfig) {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8082"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50053"
	}

	conf.Http.Port = httpPort
	conf.Grpc.Port = grpcPort
}

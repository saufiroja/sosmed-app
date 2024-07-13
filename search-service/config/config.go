package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env string
	}
	Http struct {
		Port string
	}
	Grpc struct {
		Port string
	}
	Elasticsearch struct {
		Host string
		Port string
	}
	KafkaClient struct {
		URL string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	_ = godotenv.Load("../.env")

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initServer(appConfig)
		initElasticsearch(appConfig)
		initKafkaClient(appConfig)
	}

	return appConfig
}

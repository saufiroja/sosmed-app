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
	Postgres struct {
		Name string
		User string
		Pass string
		Host string
		Port string
		SSL  string
	}
	Jwt struct {
		Secret string
	}
	AccountService struct {
		URL string
	}
	GoogleAuth struct {
		GoogleOauthClientId     string
		GoogleOauthClientSecret string
		GoogleOauthRedirectUrl  string
	}
	KafkaBroker struct {
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
		initHttp(appConfig)
		initPostgres(appConfig)
		initGoogleAuth(appConfig)
		initKafkaClient(appConfig)
	}

	return appConfig
}

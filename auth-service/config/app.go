package config

import (
	"log"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func initApp(conf *AppConfig) {
	env := os.Getenv("GO_ENV")
	secret := os.Getenv("JWT_SECRET")
	switch cases.Lower(language.English).String(env) {
	case "development":
		conf.App.Env = "development"
		conf.Jwt.Secret = secret
		log.Println("App environment is set to development")
	case "staging":
		conf.App.Env = "staging"
		conf.Jwt.Secret = secret
		log.Println("App environment is set to staging")
	case "testing":
		conf.App.Env = "testing"
		conf.Jwt.Secret = secret
		log.Println("App environment is set to testing")
	case "production":
		conf.App.Env = "production"
		conf.Jwt.Secret = secret
		log.Println("App environment is set to production")
	default:
		conf.App.Env = "development"
		conf.Jwt.Secret = secret
		log.Println("App environment is not set. Using default environment development")
	}

	conf.App.Env = env
}

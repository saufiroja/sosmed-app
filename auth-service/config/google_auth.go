package config

import "os"

func initGoogleAuth(conf *AppConfig) {
	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	redirectUrl := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

	conf.GoogleAuth.GoogleOauthClientId = clientId
	conf.GoogleAuth.GoogleOauthClientSecret = clientSecret
	conf.GoogleAuth.GoogleOauthRedirectUrl = redirectUrl
}

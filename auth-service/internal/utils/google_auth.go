package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/saufiroja/sosmed-app/auth-service/config"
	"github.com/saufiroja/sosmed-app/auth-service/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuth struct {
	configGoogle *oauth2.Config
}

func NewGoogleAuth(conf *config.AppConfig) *GoogleAuth {
	return &GoogleAuth{
		configGoogle: &oauth2.Config{
			ClientID:     conf.GoogleAuth.GoogleOauthClientId,
			ClientSecret: conf.GoogleAuth.GoogleOauthClientSecret,
			RedirectURL:  conf.GoogleAuth.GoogleOauthRedirectUrl,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (g *GoogleAuth) GoogleAuthURL() string {
	return g.configGoogle.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (g *GoogleAuth) GoogleCallback(code string) (*models.GoogleUser, error) {
	token, err := g.configGoogle.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := g.configGoogle.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var user models.GoogleUser
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return &user, nil
}

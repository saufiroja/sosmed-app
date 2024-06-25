package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saufiroja/sosmed-app/auth-service/config"
)

type GenerateToken struct {
	conf *config.AppConfig
}

func NewGenerateToken(conf *config.AppConfig) *GenerateToken {
	return &GenerateToken{
		conf: conf,
	}
}

func (g *GenerateToken) GenerateAccessToken(userId, accountType *string) (*string, error) {
	secret := g.conf.Jwt.Secret

	// 10 minutes
	exp := time.Now().Add(time.Minute * 10).Unix()
	claims := jwt.MapClaims{
		"userId":      userId,
		"accountType": accountType,
		"exp":         exp,
		"issuer":      "auth-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func (g *GenerateToken) GenerateRefreshToken(userId, accountType *string) (*string, error) {
	secret := g.conf.Jwt.Secret

	// 1 minute
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{
		"userId":      userId,
		"accountType": accountType,
		"exp":         exp,
		"issuer":      "auth-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func (g *GenerateToken) VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := g.conf.Jwt.Secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (g *GenerateToken) ValidateToken(tokenString string) (*string, *string, error) {
	token, err := g.VerifyToken(tokenString)
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, err
	}

	userId := claims["userId"].(string)
	accountType := claims["accountType"].(string)

	return &userId, &accountType, nil
}

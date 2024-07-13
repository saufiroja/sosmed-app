package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/saufiroja/sosmed-app/auth-service/config"
	"github.com/saufiroja/sosmed-app/auth-service/internal/contracts/requests"
	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	internalGrpc "github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
	"github.com/saufiroja/sosmed-app/auth-service/internal/mappers"
	"github.com/saufiroja/sosmed-app/auth-service/internal/models"
	"github.com/saufiroja/sosmed-app/auth-service/internal/repositories"
	"github.com/saufiroja/sosmed-app/auth-service/internal/utils"
	"github.com/saufiroja/sosmed-app/auth-service/pkg/database"
	"github.com/saufiroja/sosmed-app/auth-service/pkg/messaging"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepository repositories.AuthRepositoryInterface
	logger         *logrus.Logger
	db             database.PostgresInterface
	generateToken  *utils.GenerateToken
	accountClient  internalGrpc.AccountServiceClient
	googleAuth     *utils.GoogleAuth
	conf           *config.AppConfig
	kafkaProducer  *messaging.KafkaProducer
}

func NewAuthService(
	authRepository repositories.AuthRepositoryInterface,
	logger *logrus.Logger,
	db database.PostgresInterface,
	generateToken *utils.GenerateToken,
	accountClient internalGrpc.AccountServiceClient,
	googleAuth *utils.GoogleAuth,
	conf *config.AppConfig,
	kafkaProducer *messaging.KafkaProducer,
) AuthServiceInterface {
	return &authService{
		authRepository: authRepository,
		logger:         logger,
		db:             db,
		generateToken:  generateToken,
		accountClient:  accountClient,
		googleAuth:     googleAuth,
		conf:           conf,
		kafkaProducer:  kafkaProducer,
	}
}

func (a *authService) Register(ctx context.Context, request *requests.RegisterRequest) error {
	a.logger.Info(fmt.Sprintf("Registering user with account type: %s", request.AccountType))

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		UserID:        ulid.Make().String(),
		AccountTypeID: &request.AccountType,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	tx := a.db.StartTransaction(ctx)

	a.logger.Info(fmt.Sprintf("Inserting user: %s", user.UserID))
	err = a.authRepository.InsertUser(ctx, tx, user)
	if err != nil {
		a.db.RollbackTransaction(tx)
		return errors.New("failed to insert user")
	}

	insertUser := &internalGrpc.InsertUserRequest{
		UserId:   user.UserID,
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
		Password: string(hash),
	}

	_, err = a.accountClient.InsertUser(ctx, insertUser)
	if err != nil {
		a.db.RollbackTransaction(tx)
		return errors.New("failed to insert user to account service")
	}

	a.db.CommitTransaction(tx)

	// struct to string
	userByt, err := json.Marshal(insertUser)
	if err != nil {
		return err
	}

	// publish to kafka
	a.kafkaProducer.Publish(messaging.InsertUserTopic, userByt)

	return nil
}

func (a *authService) GetAccountType(ctx context.Context, accountType string) (*models.AccountType, error) {
	a.logger.Info(fmt.Sprintf("Get account type: %s", accountType))
	return a.authRepository.GetAccountType(ctx, accountType)
}

func (a *authService) Login(ctx context.Context, request *requests.LoginRequest) (*grpc.Response, error) {
	a.logger.Info(fmt.Sprintf("Login user with email: %s with account type: %s", request.Email, request.AccountType))
	input := &internalGrpc.GetAccountByEmailAndUsernameRequest{
		Email:    request.Email,
		Username: request.Username,
	}

	user, err := a.accountClient.GetAccountByEmailAndUsername(ctx, input)
	if err != nil {
		return nil, errors.New("user not found")
	}

	accountType, err := a.GetAccountType(ctx, request.AccountType)
	if err != nil {
		return nil, errors.New("account type not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// access token
	accessToken, err := a.generateToken.GenerateAccessToken(&user.UserId, &accountType.AccountName)
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshToken, err := a.generateToken.GenerateRefreshToken(&user.UserId, &accountType.AccountName)
	if err != nil {
		return nil, err
	}

	response := mappers.ToTokenResponse(accessToken, refreshToken)

	return response, nil
}

func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (*grpc.Response, error) {
	userId, accountType, err := a.generateToken.ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// access token
	newAccessToken, err := a.generateToken.GenerateAccessToken(userId, accountType)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	// refresh token
	newRefreshToken, err := a.generateToken.GenerateRefreshToken(userId, accountType)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	response := mappers.ToTokenResponse(newAccessToken, newRefreshToken)

	return response, nil
}

func (a *authService) GoogleAuth(ctx context.Context) (*string, error) {
	a.logger.Info("Google Auth")

	url := a.googleAuth.GoogleAuthURL()
	a.logger.Info(fmt.Sprintf("Google Auth URL: %s", url))

	return &url, nil
}

func (a *authService) GoogleAuthCallback(ctx context.Context, code string) (*grpc.Response, error) {
	userInfo, err := a.googleAuth.GoogleCallback(code)
	if err != nil {
		return nil, err
	}

	accountType, err := a.GetAccountType(ctx, "Google")
	if err != nil {
		return nil, errors.New("account type not found")
	}

	user := &models.User{
		UserID:        ulid.Make().String(),
		AccountTypeID: &accountType.ID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	tx := a.db.StartTransaction(ctx)

	a.logger.Info(fmt.Sprintf("Inserting user: %s", user.UserID))
	err = a.authRepository.InsertUser(ctx, tx, user)
	if err != nil {
		a.db.RollbackTransaction(tx)
		return nil, errors.New("failed to insert user")
	}

	insertUser := &internalGrpc.InsertUserRequest{
		UserId:   user.UserID,
		FullName: userInfo.Name,
		Email:    userInfo.Email,
		Username: "ttestt",
	}

	_, err = a.accountClient.InsertUser(ctx, insertUser)
	if err != nil {
		a.db.RollbackTransaction(tx)
		return nil, errors.New("failed to insert user to account service")
	}

	a.db.CommitTransaction(tx)

	accessToken, err := a.generateToken.GenerateAccessToken(&user.UserID, &accountType.AccountName)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.generateToken.GenerateRefreshToken(&user.UserID, &accountType.AccountName)
	if err != nil {
		return nil, err
	}

	response := mappers.ToTokenResponse(accessToken, refreshToken)

	return response, nil
}

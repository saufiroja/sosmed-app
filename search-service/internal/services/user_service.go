package services

import (
	"context"
	"fmt"

	"github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/saufiroja/sosmed/search-service/internal/models"
	"github.com/saufiroja/sosmed/search-service/internal/repositories"
	"github.com/sirupsen/logrus"
)

type userService struct {
	userRepository repositories.UserRepositoryInterface
	logger         *logrus.Logger
}

func NewUserService(userRepository repositories.UserRepositoryInterface, logger *logrus.Logger) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (u *userService) InsertUser(ctx context.Context, user *models.User) error {
	u.logger.Info(fmt.Sprintf("inserting user: %v", user))
	return u.userRepository.InsertUser(ctx, user)
}

func (u *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return u.userRepository.GetUserByUsername(ctx, username)
}

func (u *userService) GetAllUsers(ctx context.Context, request *grpc.SearchAllUsersRequest) ([]models.User, error) {
	u.logger.Info(fmt.Sprintf("getting all users from %v to %v", request.Page, request.Limit))

	return u.userRepository.GetAllUsers(ctx, &request.Page, &request.Limit)
}

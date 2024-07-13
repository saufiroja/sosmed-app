package query

import (
	"context"

	"github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/saufiroja/sosmed/search-service/internal/services"
)

type GetAllUsersQueryHandler struct {
	userService services.UserServiceInterface
}

func NewGetAllUsersQueryHandler(userService services.UserServiceInterface) *GetAllUsersQueryHandler {
	return &GetAllUsersQueryHandler{userService: userService}
}
func (h *GetAllUsersQueryHandler) Handle(ctx context.Context, request *grpc.SearchAllUsersRequest) (*grpc.SearchAllUsersResponse, error) {
	res, err := h.userService.GetAllUsers(ctx, request)
	if err != nil {
		return nil, err
	}

	var users []*grpc.GetAllUsers
	for _, user := range res {
		users = append(users, &grpc.GetAllUsers{
			UserId:   user.UserID,
			Username: user.Username,
			FullName: user.FullName,
		})
	}

	return &grpc.SearchAllUsersResponse{
		Users: users,
	}, nil
}

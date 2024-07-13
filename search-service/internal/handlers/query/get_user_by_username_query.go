package query

import (
	"context"

	"github.com/saufiroja/sosmed/search-service/internal/grpc"
	"github.com/saufiroja/sosmed/search-service/internal/services"
)

type GetUserByUsernameQueryHandler struct {
	userService services.UserServiceInterface
}

func NewGetUserByUsernameQueryHandler(userService services.UserServiceInterface) *GetUserByUsernameQueryHandler {
	return &GetUserByUsernameQueryHandler{userService: userService}
}
func (h *GetUserByUsernameQueryHandler) Handle(ctx context.Context, request *grpc.SearchUserByUsernameRequest) (*grpc.SearchUserByUsernameResponse, error) {
	res, err := h.userService.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	return &grpc.SearchUserByUsernameResponse{
		Username: res.Username,
	}, nil
}

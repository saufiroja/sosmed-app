package mappers

import (
	"github.com/saufiroja/sosmed-app/auth-service/internal/grpc"
)

func ToTokenResponse(accessToken, refreshToken *string) *grpc.Response {
	return &grpc.Response{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}
}

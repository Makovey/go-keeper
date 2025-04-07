package mock

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/gen/auth"
)

type authClientMock struct {
	model *auth.AuthResponse
	error error
}

func NewAuthClientMock(
	model *auth.AuthResponse,
	error error,
) auth.AuthClient {
	return &authClientMock{
		model: model,
		error: error,
	}
}

func (a authClientMock) RegisterUser(ctx context.Context, in *auth.User, opts ...grpc.CallOption) (*auth.AuthResponse, error) {
	if a.error != nil {
		return nil, a.error
	}

	return a.model, nil
}

func (a authClientMock) LoginUser(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.AuthResponse, error) {
	if a.error != nil {
		return nil, a.error
	}

	return a.model, nil
}

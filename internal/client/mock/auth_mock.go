package mock

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/gen/auth"
)

type authClientMock struct {
	registerResponse *auth.AuthResponse
	registerError    error
	loginResponse    *auth.AuthResponse
	loginUserError   error
}

func NewAuthClientMock(
	registerResponse *auth.AuthResponse,
	registerError error,
	loginResponse *auth.AuthResponse,
	loginUserError error,
) auth.AuthClient {
	return &authClientMock{
		registerResponse: registerResponse,
		registerError:    registerError,
		loginResponse:    loginResponse,
		loginUserError:   loginUserError,
	}
}

func (a authClientMock) RegisterUser(ctx context.Context, in *auth.User, opts ...grpc.CallOption) (*auth.AuthResponse, error) {
	if a.registerError != nil {
		return nil, a.registerError
	}

	return a.registerResponse, nil
}

func (a authClientMock) LoginUser(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.AuthResponse, error) {
	if a.loginUserError != nil {
		return nil, a.loginUserError
	}

	return a.loginResponse, nil
}

package mock

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/Makovey/go-keeper/internal/gen/auth"
)

type authClientMock struct {
	model *pb.AuthResponse
	error error
}

func NewAuthClientMock(
	model *pb.AuthResponse,
	error error,
) pb.AuthClient {
	return &authClientMock{
		model: model,
		error: error,
	}
}

func (a authClientMock) RegisterUser(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	if a.error != nil {
		return nil, a.error
	}

	return a.model, nil
}

func (a authClientMock) LoginUser(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	if a.error != nil {
		return nil, a.error
	}

	return a.model, nil
}

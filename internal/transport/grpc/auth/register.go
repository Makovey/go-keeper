package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/service"
	grpcerr "github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
)

func (s *Server) RegisterUser(ctx context.Context, protoUser *auth.User) (*emptypb.Empty, error) {
	fn := "auth.RegisterUser"

	user := mapper.ToUserFromProto(protoUser)

	if err := user.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.service.RegisterUser(ctx, user)
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err.Error())
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, "email already registered")
		}
		return nil, status.Error(codes.Internal, grpcerr.InternalServerError)
	}

	md := metadata.New(map[string]string{"jwt": token})
	if err = grpc.SetHeader(ctx, md); err != nil {
		s.log.Errorf("[%s]: %v", fn, err.Error())
		return nil, status.Error(codes.Internal, grpcerr.InternalServerError)
	}

	return nil, nil
}

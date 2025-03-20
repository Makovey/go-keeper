package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/service"
	"github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
)

func (s *Server) RegisterUser(ctx context.Context, user *auth.User) (*emptypb.Empty, error) {
	fn := "auth.RegisterUser"

	err := s.service.RegisterUser(ctx, mapper.ToUserFromProto(user))
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err.Error())
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, "email already registered")
		}
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	return nil, nil
}

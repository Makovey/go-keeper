package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/service"
	grpcerr "github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	fn := "auth.LoginUser"

	login := mapper.ToLoginFromProto(req)

	if err := login.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.service.LoginUser(ctx, login)
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err.Error())
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			return nil, status.Error(codes.InvalidArgument, "email not found")
		case errors.Is(err, service.ErrIncorrectPassword):
			return nil, status.Error(codes.InvalidArgument, "incorrect password")
		}
		return nil, status.Error(codes.Internal, grpcerr.InternalServerError)
	}

	md := metadata.New(map[string]string{"jwt": token})
	if err = grpc.SetHeader(ctx, md); err != nil {
		s.log.Errorf("[%s]: %v", fn, err.Error())
		return nil, status.Error(codes.Internal, grpcerr.InternalServerError)
	}

	return &pb.AuthResponse{Token: token}, nil
}

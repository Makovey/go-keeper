package storage

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
)

func (s *Server) GetUsersFile(ctx context.Context, req *emptypb.Empty) (*storage.GetUsersFileResponse, error) {
	fn := "storage.GetUsersFile"

	userID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	files, err := s.service.GetUsersFiles(ctx, userID)
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return nil, status.Error(codes.Internal, helper.InternalServerError)
	}

	return &storage.GetUsersFileResponse{Files: mapper.ToProtoFromFile(files)}, nil
}

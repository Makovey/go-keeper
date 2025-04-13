package storage

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
)

func (s *Server) DeleteUsersFile(ctx context.Context, req *pb.DeleteUsersFileRequest) (*emptypb.Empty, error) {
	fn := "storage.DeleteUsersFile"

	userID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	if err = s.service.DeleteUsersFile(ctx, userID, req.GetFileId(), req.GetFileName()); err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return nil, status.Error(codes.Internal, helper.InternalServerError)
	}

	return &emptypb.Empty{}, nil
}

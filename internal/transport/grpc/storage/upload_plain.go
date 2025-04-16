package storage

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
)

func (s *Server) UploadPlainTextType(ctx context.Context, req *pb.UploadPlainTextTypeRequest) (*pb.UploadPlainTextTypeResponse, error) {
	fn := "storage.UploadFile"

	userID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	name, err := s.service.UploadPlainText(ctx, userID, req.GetContent())
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return nil, status.Error(codes.Internal, helper.InternalServerError)
	}

	return &pb.UploadPlainTextTypeResponse{FileName: name}, nil
}

package storage

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
)

func (s *Server) UploadPlainTextType(ctx context.Context, req *storage.UploadPlainTextTypeRequest) (*storage.UploadPlainTextTypeResponse, error) {
	fn := "storage.UploadFile"

	userID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	name, err := s.service.UploadPlainText(ctx, userID, req.GetContent(), helper.MapProtoToLocalTextSecure(req.GetType()))
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return nil, status.Error(codes.Internal, helper.InternalServerError)
	}

	return &storage.UploadPlainTextTypeResponse{FileName: name}, nil
}

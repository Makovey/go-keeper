package storage

import (
	"context"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

//go:generate mockgen -source=server.go -destination=../../../service/mock/storage_service_mock.go -package=mock
type ServiceStorage interface {
	UploadFile(ctx context.Context, file model.File, userID string) (string, error)
	DownloadFile(ctx context.Context, userID, fileID string) (*model.File, error)
	GetUsersFiles(ctx context.Context, userID string) ([]*model.ExtendedInfoFile, error)
	DeleteUsersFile(ctx context.Context, userID, fileID, fileName string) error
}

type Server struct {
	storage.UnimplementedStorageServiceServer

	log     logger.Logger
	service ServiceStorage
}

func NewStorageServer(
	log logger.Logger,
	service ServiceStorage,
) *Server {
	return &Server{
		log:     log,
		service: service,
	}
}

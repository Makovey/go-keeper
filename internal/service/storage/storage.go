package storage

import (
	"context"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/repository/entity"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
	"github.com/Makovey/go-keeper/internal/transport/grpc/storage"
)

//go:generate mockgen -source=storage.go -destination=../../repository/mock/storage_repository_mock.go -package=mock
type Repository interface {
	SaveFileMetadata(ctx context.Context, fileData *entity.File) error
}

type service struct {
	repo Repository
	cfg  config.Config
	log  logger.Logger
}

func NewStorageService(
	repo Repository,
	cfg config.Config,
	log logger.Logger,
) storage.Service {
	return &service{
		repo: repo,
		cfg:  cfg,
		log:  log,
	}
}

func (s *service) UploadFile(ctx context.Context, file model.File, userID string) (string, error) {
	// TODO: add impl
	return "", nil
}

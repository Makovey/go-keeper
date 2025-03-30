package storage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/repository/entity"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
	"github.com/Makovey/go-keeper/internal/transport/grpc/storage"
)

type FileStorager interface {
	Save(path, fileName string, data bytes.Reader) error
}

//go:generate mockgen -source=storage.go -destination=../../repository/mock/storage_repository_mock.go -package=mock
type Repository interface {
	SaveFileMetadata(ctx context.Context, fileData *entity.File) error
}

type service struct {
	repo     Repository
	storager FileStorager
	cfg      config.Config
	log      logger.Logger
}

func NewStorageService(
	repo Repository,
	storager FileStorager,
	log logger.Logger,
) storage.Service {
	return &service{
		repo:     repo,
		storager: storager,
		log:      log,
	}
}

func (s *service) UploadFile(ctx context.Context, file model.File, userID string) (string, error) {
	fn := "storage.UploadFile"

	eFile := &entity.File{
		ID:       uuid.NewString(),
		OwnerID:  userID,
		FileName: file.FileName,
		FileSize: formatFileSize(file.FileSize),
		Path:     fmt.Sprintf("%s/%s", userID, file.FileName),
	}

	if err := s.storager.Save(userID, file.FileName, file.Data); err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	if err := s.repo.SaveFileMetadata(ctx, eFile); err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	return eFile.ID, nil
}

func formatFileSize(bytes int) string {
	const (
		KB = 1000
		MB = KB * 1000
		GB = MB * 1000
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.1f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.1f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.1f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

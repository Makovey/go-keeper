package storage

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/repository/entity"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
	"github.com/Makovey/go-keeper/internal/transport/grpc/storage"
)

type FileStorager interface {
	Save(path, fileName string, data *bufio.Reader) error
	Get(path string) ([]byte, error)
}

// RepositoryStorage NOTE: префикс Storage, чтобы не было коллизии имен при генерации моков
//
//go:generate mockgen -source=storage.go -destination=../../repository/mock/storage_repository_mock.go -package=mock
type RepositoryStorage interface {
	SaveFileMetadata(ctx context.Context, fileData *entity.File) error
	GetFileMetadata(ctx context.Context, userID, fileID string) (*entity.File, error)
	GetUsersFiles(ctx context.Context, userID string) ([]*entity.File, error)
}

type service struct {
	repo     RepositoryStorage
	storager FileStorager
	cfg      config.Config
}

func NewStorageService(
	repo RepositoryStorage,
	storager FileStorager,
) storage.ServiceStorage {
	return &service{
		repo:     repo,
		storager: storager,
	}
}

func (s *service) UploadFile(ctx context.Context, file model.File, userID string) (string, error) {
	fn := "storage.UploadFile"

	eFile := &entity.File{
		ID:       uuid.NewString(),
		OwnerID:  userID,
		FileName: file.FileName,
		FileSize: file.FileSize,
		Path:     fmt.Sprintf("%s/%s", userID, file.FileName),
	}

	if err := s.storager.Save(userID, file.FileName, &file.Data); err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	if err := s.repo.SaveFileMetadata(ctx, eFile); err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	return eFile.ID, nil
}

func (s *service) DownloadFile(ctx context.Context, userID, fileID string) (*model.File, error) {
	fn := "storage.DownloadFile"

	file, err := s.repo.GetFileMetadata(ctx, userID, fileID)
	if err != nil {
		return &model.File{}, fmt.Errorf("[%s]: %v", fn, err)
	}

	data, err := s.storager.Get(file.Path)
	if err != nil {
		return &model.File{}, fmt.Errorf("[%s]: %v", fn, err)
	}

	return &model.File{Data: *bufio.NewReader(bytes.NewReader(data)), FileName: file.FileName, FileSize: file.FileSize}, nil
}

func (s *service) GetUsersFiles(ctx context.Context, userID string) ([]*model.ExtendedInfoFile, error) {
	fn := "storage.GetUsersFiles"

	files, err := s.repo.GetUsersFiles(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %v", fn, err)
	}

	res := make([]*model.ExtendedInfoFile, 0, len(files))
	for _, file := range files {
		res = append(res, &model.ExtendedInfoFile{
			ID:        file.ID,
			FileName:  file.FileName,
			FileSize:  formatFileSize(file.FileSize),
			CreatedAt: file.CreatedAt,
		})
	}

	return res, nil
}

func formatFileSize(bytes int) string {
	switch {
	case bytes >= humanize.GByte:
		return fmt.Sprintf("%.1f GB", float64(bytes)/humanize.GByte)
	case bytes >= humanize.MByte:
		return fmt.Sprintf("%.1f MB", float64(bytes)/humanize.MByte)
	case bytes >= humanize.KByte:
		return fmt.Sprintf("%.1f KB", float64(bytes)/humanize.KByte)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

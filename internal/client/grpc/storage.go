package grpc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/utils"
)

const (
	rootDir = "go-keeper"
)

type StorageClient struct {
	log    logger.Logger
	client storage.StorageServiceClient
}

func NewStorageClient(
	log logger.Logger,
	client storage.StorageServiceClient,
) *StorageClient {

	return &StorageClient{
		log:    log,
		client: client,
	}
}

func (s *StorageClient) UploadFile(
	ctx context.Context,
	path string,
) error {
	fn := "grpc.UploadFile"

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("[%s]: failed to open file: %v", fn, err)
	}
	defer file.Close()

	stream, err := s.client.UploadFile(ctx)
	if err != nil {
		return err
	}

	if err = stream.Send(&storage.UploadRequest{
		FileName: filepath.Base(path),
	}); err != nil {
		return fmt.Errorf("[%s]: failed to send request: %v", fn, err)
	}

	buf := make([]byte, humanize.MByte)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("[%s]: failed to read file: %v", fn, err)
		}

		if err := stream.Send(&storage.UploadRequest{
			ChunkData: buf[:n],
		}); err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("[%s]: failed to close stream: %v", fn, err)
	}

	return nil
}

func (s *StorageClient) DownloadFile(
	ctx context.Context,
	fileID string,
) error {
	fn := "grpc.DownloadFile"

	req := &storage.DownloadRequest{FileId: fileID}
	stream, err := s.client.DownloadFile(ctx, req)
	if err != nil {
		return fmt.Errorf("[%s]: failed to init download: %v", fn, err)
	}

	if err = utils.CreateDirIfNeeded(rootDir, fileID); err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	fullPath := fmt.Sprintf("./%s/%s", rootDir, fileID)
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("[%s]: failed to create file: %v", fn, err)
	}
	defer file.Close()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("[%s]: download failed: %v", fn, err)
		}

		if _, err := file.Write(res.ChunkData); err != nil {
			return fmt.Errorf("[%s]: failed to write chunk: %v", fn, err)
		}
	}

	return nil
}

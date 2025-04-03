package grpc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger"
)

type StorageClient struct {
	log    logger.Logger
	client storage.StorageServiceClient
}

func NewStorageClient(
	conn *grpc.ClientConn,
	log logger.Logger,
) *StorageClient {

	return &StorageClient{
		log:    log,
		client: storage.NewStorageServiceClient(conn),
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
		Filename: filepath.Base(path),
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

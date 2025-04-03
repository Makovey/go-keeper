package grpc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger"
)

type StorageClient struct {
	cfg    config.Config
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

	md := metadata.New(map[string]string{"jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM3MDgzMjgsIlVzZXJJRCI6ImM5MDllMTdkLTg0MzMtNGI4ZC05ZDE2LTFiZmY2NzVmNGEzNiJ9.5U2UHDDSfzVxjPnN4sCkuHTrf1jllmPSF4EgfSP2tH4"})
	ctx = metadata.NewOutgoingContext(ctx, md)

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

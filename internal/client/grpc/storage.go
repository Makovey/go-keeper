package grpc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/utils"
)

const (
	rootDir = "go-keeper"
	dirName = "files"
)

type StorageClient struct {
	log        logger.Logger
	dirManager utils.DirManager
	client     pb.StorageServiceClient
}

func NewStorageClient(
	log logger.Logger,
	dirManager utils.DirManager,
	client pb.StorageServiceClient,
) *StorageClient {

	return &StorageClient{
		log:        log,
		dirManager: dirManager,
		client:     client,
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

	if err = stream.Send(&pb.UploadRequest{
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

		if err := stream.Send(&pb.UploadRequest{
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

func (s *StorageClient) GetUsersFiles(
	ctx context.Context,
) ([]*pb.UsersFile, error) {
	fn := "grpc.GetUserFiles"

	resp, err := s.client.GetUsersFile(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("[%s]: failed to get users files: %v", fn, err)
	}

	return resp.GetFiles(), nil
}

func (s *StorageClient) DownloadFile(
	ctx context.Context,
	fileID string,
) error {
	fn := "grpc.DownloadFile"

	req := &pb.DownloadRequest{FileId: fileID}
	stream, err := s.client.DownloadFile(ctx, req)
	if err != nil {
		return fmt.Errorf("[%s]: failed to init download: %v", fn, err)
	}

	firstChunk, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("[%s]: failed to get filename: %v", fn, err)
	}

	if firstChunk.GetFileName() == "" {
		return fmt.Errorf("[%s]: empty filename received", fn)
	}

	if err = s.dirManager.CreateDir(rootDir, dirName); err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	fullPath := fmt.Sprintf("./%s/%s/%s", rootDir, dirName, firstChunk.GetFileName())
	file, err := s.dirManager.CreateFile(fullPath)
	if err != nil {
		return fmt.Errorf("[%s]: failed to create file: %v", fn, err)
	}
	defer file.Close()

	if _, err := file.Write(firstChunk.ChunkData); err != nil {
		return fmt.Errorf("[%s]: failed to write first chunk: %v", fn, err)
	}

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

func (s *StorageClient) DeleteFile(
	ctx context.Context,
	fileID string,
	fileName string,
) error {
	fn := "grpc.DeleteFile"

	req := &pb.DeleteUsersFileRequest{FileName: fileName, FileId: fileID}
	_, err := s.client.DeleteUsersFile(ctx, req)
	if err != nil {
		return fmt.Errorf("[%s]: failed to delete file: %v", fn, err)
	}

	return nil
}

func (s *StorageClient) UploadPlainText(
	ctx context.Context,
	content string,
) (string, error) {
	fn := "grpc.DeleteFile"

	req := &pb.UploadPlainTextTypeRequest{Content: content}
	resp, err := s.client.UploadPlainTextType(ctx, req)
	if err != nil {
		return "", fmt.Errorf("[%s]: failed to upload plain text: %v", fn, err)
	}

	return resp.GetFileName(), nil
}

package mock

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/gen/storage"
)

type storageClientMock struct {
	uploadStream      grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse]
	downloadStream    grpc.ServerStreamingClient[storage.DownloadResponse]
	usersFileResponse *storage.GetUsersFileResponse
	error             error
}

func NewStorageWithUploadStream(
	stream grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse],
	error error,
) storage.StorageServiceClient {
	return &storageClientMock{
		uploadStream: stream,
		error:        error,
	}
}

func NewStorageWithDownloadStream(
	stream grpc.ServerStreamingClient[storage.DownloadResponse],
	error error,
) storage.StorageServiceClient {
	return &storageClientMock{
		downloadStream: stream,
		error:          error,
	}
}

func NewStorageWitUsersFile(
	usersFileResponse *storage.GetUsersFileResponse,
	error error,
) storage.StorageServiceClient {
	return &storageClientMock{
		usersFileResponse: usersFileResponse,
		error:             error,
	}
}

func NewStorageWitEmptyMock(
	error error,
) storage.StorageServiceClient {
	return &storageClientMock{
		error: error,
	}
}

func (s *storageClientMock) UploadFile(
	ctx context.Context,
	opts ...grpc.CallOption,
) (grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse], error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.uploadStream, nil
}

func (s *storageClientMock) GetUsersFile(
	ctx context.Context,
	in *emptypb.Empty,
	opts ...grpc.CallOption,
) (*storage.GetUsersFileResponse, error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.usersFileResponse, nil
}

func (s *storageClientMock) DownloadFile(ctx context.Context, in *storage.DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[storage.DownloadResponse], error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.downloadStream, nil
}

func (s *storageClientMock) DeleteUsersFile(ctx context.Context, in *storage.DeleteUsersFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	if s.error != nil {
		return nil, s.error
	}

	return &emptypb.Empty{}, nil
}

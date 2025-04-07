package mock

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/gen/storage"
)

type storageClientMock struct {
	uploadResponse grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse]
	uploadError    error
}

func NewStorageClientMock(
	uploadResponse grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse],
	uploadError error,
) storage.StorageServiceClient {
	return &storageClientMock{
		uploadResponse: uploadResponse,
		uploadError:    uploadError,
	}
}

func (s storageClientMock) UploadFile(
	ctx context.Context, opts ...grpc.CallOption,
) (grpc.ClientStreamingClient[storage.UploadRequest, storage.UploadResponse], error) {
	if s.uploadError != nil {
		return nil, s.uploadError
	}

	return s.uploadResponse, nil
}

func (s storageClientMock) DownloadFile(ctx context.Context, in *storage.DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[storage.DownloadResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s storageClientMock) GetUsersFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*storage.GetUsersFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageClientMock) DeleteUsersFile(ctx context.Context, in *storage.DeleteUsersFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

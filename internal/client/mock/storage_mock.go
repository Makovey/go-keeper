package mock

import (
	"context"

	"google.golang.org/grpc"

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

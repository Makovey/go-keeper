package mock

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
)

type storageClientMock struct {
	uploadStream        grpc.ClientStreamingClient[pb.UploadRequest, pb.UploadResponse]
	downloadStream      grpc.ServerStreamingClient[pb.DownloadResponse]
	usersFileResponse   *pb.GetUsersFileResponse
	uploadPlainResponse *pb.UploadPlainTextTypeResponse
	error               error
}

func NewStorageWithUploadStream(
	stream grpc.ClientStreamingClient[pb.UploadRequest, pb.UploadResponse],
	error error,
) pb.StorageServiceClient {
	return &storageClientMock{
		uploadStream: stream,
		error:        error,
	}
}

func NewStorageWithDownloadStream(
	stream grpc.ServerStreamingClient[pb.DownloadResponse],
	error error,
) pb.StorageServiceClient {
	return &storageClientMock{
		downloadStream: stream,
		error:          error,
	}
}

func NewStorageWitUsersFile(
	usersFileResponse *pb.GetUsersFileResponse,
	error error,
) pb.StorageServiceClient {
	return &storageClientMock{
		usersFileResponse: usersFileResponse,
		error:             error,
	}
}

func NewStorageWitUploadPlainText(
	uploadResponse *pb.UploadPlainTextTypeResponse,
	error error,
) pb.StorageServiceClient {
	return &storageClientMock{
		uploadPlainResponse: uploadResponse,
		error:               error,
	}
}

func NewStorageWitEmptyMock(
	error error,
) pb.StorageServiceClient {
	return &storageClientMock{
		error: error,
	}
}

func (s *storageClientMock) UploadFile(
	ctx context.Context,
	opts ...grpc.CallOption,
) (grpc.ClientStreamingClient[pb.UploadRequest, pb.UploadResponse], error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.uploadStream, nil
}

func (s *storageClientMock) GetUsersFile(
	ctx context.Context,
	in *emptypb.Empty,
	opts ...grpc.CallOption,
) (*pb.GetUsersFileResponse, error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.usersFileResponse, nil
}

func (s *storageClientMock) DownloadFile(ctx context.Context, in *pb.DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.DownloadResponse], error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.downloadStream, nil
}

func (s *storageClientMock) DeleteUsersFile(ctx context.Context, in *pb.DeleteUsersFileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	if s.error != nil {
		return nil, s.error
	}

	return &emptypb.Empty{}, nil
}

func (s *storageClientMock) UploadPlainTextType(
	ctx context.Context,
	in *pb.UploadPlainTextTypeRequest,
	opts ...grpc.CallOption,
) (*pb.UploadPlainTextTypeResponse, error) {
	if s.error != nil {
		return nil, s.error
	}

	return s.uploadPlainResponse, nil
}

package mock

import (
	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/gen/storage"
)

type ClientStreamMock struct {
	grpc.ClientStream
	SendFunc         func(*storage.UploadRequest) error
	CloseAndRecvFunc func() (*storage.UploadResponse, error)
}

func (m *ClientStreamMock) Send(req *storage.UploadRequest) error {
	return m.SendFunc(req)
}

func (m *ClientStreamMock) CloseAndRecv() (*storage.UploadResponse, error) {
	return m.CloseAndRecvFunc()
}

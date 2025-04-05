package mock

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/gen/storage"
)

type ClientStreamMock struct {
	grpc.ClientStreamingServer[storage.UploadRequest, storage.UploadResponse]
	RecvFunc         func() (*storage.UploadRequest, error)
	SendAndCloseFunc func(*storage.UploadResponse) error
	ContextFunc      func() context.Context
}

func (m *ClientStreamMock) Recv() (*storage.UploadRequest, error) {
	return m.RecvFunc()
}

func (m *ClientStreamMock) SendAndClose(resp *storage.UploadResponse) error {
	return m.SendAndCloseFunc(resp)
}

func (m *ClientStreamMock) Context() context.Context {
	return m.ContextFunc()
}

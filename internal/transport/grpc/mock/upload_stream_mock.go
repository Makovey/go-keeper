package mock

import (
	"context"

	"google.golang.org/grpc"
)

type ClientStreamMock[T any, I any] struct {
	grpc.ClientStreamingServer[T, I]
	RecvFunc         func() (*T, error)
	SendAndCloseFunc func(*I) error
	ContextFunc      func() context.Context
	SendFunc         func(*I) error
}

func (m *ClientStreamMock[T, I]) Recv() (*T, error) {
	return m.RecvFunc()
}

func (m *ClientStreamMock[T, I]) SendAndClose(resp *I) error {
	return m.SendAndCloseFunc(resp)
}

func (m *ClientStreamMock[T, I]) Context() context.Context {
	return m.ContextFunc()
}

func (m *ClientStreamMock[T, I]) Send(res *I) error {
	return m.SendFunc(res)
}

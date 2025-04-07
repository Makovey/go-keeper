package mock

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientStreamMock[T any, I any] struct {
	grpc.ClientStream
	SendFunc         func(*T) error
	CloseAndRecvFunc func() (*I, error)
}

func (m *ClientStreamMock[T, I]) Send(req *T) error {
	return m.SendFunc(req)
}

func (m *ClientStreamMock[T, I]) CloseAndRecv() (*I, error) {
	return m.CloseAndRecvFunc()
}

type ServerStreamClientMock[T any] struct {
	grpc.ServerStream
	RecvFunc         func() (*T, error)
	CloseAndRecvFunc func() (*T, error)
	HeaderFunc       func() (metadata.MD, error)
	TrailerFunc      func() metadata.MD
	CloseSendFunc    func() error
}

func (s *ServerStreamClientMock[T]) Recv() (*T, error) {
	return s.RecvFunc()
}

func (s *ServerStreamClientMock[T]) CloseAndRecv() (*T, error) {
	return s.CloseAndRecvFunc()
}

func (s *ServerStreamClientMock[T]) Header() (metadata.MD, error) {
	return s.HeaderFunc()
}

func (s *ServerStreamClientMock[T]) Trailer() metadata.MD {
	return s.TrailerFunc()
}

func (s *ServerStreamClientMock[T]) CloseSend() error {
	return s.CloseSendFunc()
}

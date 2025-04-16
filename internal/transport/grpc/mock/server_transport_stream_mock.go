package mock

import "google.golang.org/grpc/metadata"

type ServerTransportStreamMock struct {
	md metadata.MD
}

func (m *ServerTransportStreamMock) Method() string {
	return ""
}

func (m *ServerTransportStreamMock) SetHeader(md metadata.MD) error {
	m.md = md
	return nil
}

func (m *ServerTransportStreamMock) SendHeader(md metadata.MD) error {
	return nil
}

func (m *ServerTransportStreamMock) SetTrailer(md metadata.MD) error {
	return nil
}

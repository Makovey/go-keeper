package storage

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/service/mock"
	grpcMock "github.com/Makovey/go-keeper/internal/transport/grpc/mock"
)

func TestServer_UploadFile(t *testing.T) {
	type args struct {
		jwtToken string
		chunks   []*storage.UploadRequest
	}

	type expects struct {
		servErr error
		servAns string
		wantErr bool
		result  codes.Code
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "successfully upload file",
			args: args{
				jwtToken: uuid.NewString(),
				chunks: []*storage.UploadRequest{
					{
						Filename:  "testable.txt",
						ChunkData: []byte("test1"),
					},
					{
						ChunkData: []byte("test2"),
					},
				}},
			expects: expects{servAns: uuid.NewString(), result: codes.OK},
		},
		{
			name: "failed to upload file: invalid jwtToken",
			args: args{
				jwtToken: "testToken",
				chunks: []*storage.UploadRequest{
					{
						Filename:  "testable.txt",
						ChunkData: []byte("test1"),
					},
					{
						ChunkData: []byte("test2"),
					},
				}},
			expects: expects{servAns: uuid.NewString(), result: codes.Unauthenticated, wantErr: true},
		},
		{
			name: "failed to upload file: invalid request",
			args: args{
				jwtToken: uuid.NewString(),
				chunks:   []*storage.UploadRequest{},
			},
			expects: expects{servAns: uuid.NewString(), result: codes.InvalidArgument, wantErr: true},
		},
		{
			name: "failed to upload file: service error",
			args: args{
				jwtToken: uuid.NewString(),
				chunks: []*storage.UploadRequest{
					{
						Filename:  "testable.txt",
						ChunkData: []byte("test1"),
					},
					{
						ChunkData: []byte("test2"),
					},
				}},
			expects: expects{
				servErr: errors.New("something went wrong"),
				result:  codes.Internal,
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), jwt.CtxUserIDKey, tt.args.jwtToken)
			m := mock.NewMockServiceStorage(ctrl)
			m.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.servAns, tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			stream := &grpcMock.ClientStreamMock{
				RecvFunc: func() func() (*storage.UploadRequest, error) {
					count := 0
					return func() (*storage.UploadRequest, error) {
						if len(tt.args.chunks) < 1 {
							return nil, errors.New("test error")
						}
						count++
						switch count {
						case 1:
							return tt.args.chunks[count-1], nil
						case 2:
							return tt.args.chunks[count-1], nil
						default:
							return nil, io.EOF
						}
					}
				}(),
				SendAndCloseFunc: func(resp *storage.UploadResponse) error {
					assert.Equal(t, tt.expects.servAns, resp.FileId)
					return nil
				},
				ContextFunc: func() context.Context {
					return ctx
				},
			}

			err := s.UploadFile(stream)
			if tt.expects.wantErr {
				assert.Equal(t, tt.expects.result, status.Code(err))
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expects.result, codes.OK)
				assert.NoError(t, err)
			}
		})
	}
}

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

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/service/mock"
	grpcMock "github.com/Makovey/go-keeper/internal/transport/grpc/mock"
)

func TestServer_UploadFile(t *testing.T) {
	type args struct {
		userID string
		chunks []*pb.UploadRequest
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
				userID: uuid.NewString(),
				chunks: []*pb.UploadRequest{
					{
						FileName:  "testable.txt",
						ChunkData: []byte("test1"),
					},
					{
						ChunkData: []byte("test2"),
					},
				}},
			expects: expects{servAns: uuid.NewString(), result: codes.OK},
		},
		{
			name: "failed to upload file: invalid userID",
			args: args{
				userID: "testToken",
				chunks: []*pb.UploadRequest{
					{
						FileName:  "testable.txt",
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
				userID: uuid.NewString(),
				chunks: []*pb.UploadRequest{},
			},
			expects: expects{servAns: uuid.NewString(), result: codes.Internal, wantErr: true},
		},
		{
			name: "failed to upload file: service error",
			args: args{
				userID: uuid.NewString(),
				chunks: []*pb.UploadRequest{
					{
						FileName:  "testable.txt",
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

			ctx := context.WithValue(context.Background(), jwt.CtxUserIDKey, tt.args.userID)
			m := mock.NewMockServiceStorage(ctrl)
			m.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.servAns, tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			stream := &grpcMock.ClientStreamMock[pb.UploadRequest, pb.UploadResponse]{
				RecvFunc: func() func() (*pb.UploadRequest, error) {
					count := 0
					return func() (*pb.UploadRequest, error) {
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
				SendAndCloseFunc: func(resp *pb.UploadResponse) error {
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

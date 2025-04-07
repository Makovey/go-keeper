package storage

import (
	"bufio"
	"bytes"
	"context"
	"errors"
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
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func TestServer_DownloadFile(t *testing.T) {
	type args struct {
		userID string
		fileID string
	}

	type expects struct {
		servErr error
		servAns *model.File
		sendErr error
		wantErr bool
		result  codes.Code
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "successfully download file",
			args: args{
				userID: uuid.NewString(),
			},
			expects: expects{
				servAns: &model.File{
					FileName: "testfile.txt",
					FileSize: 100,
					Data:     *bufio.NewReader(bytes.NewReader([]byte("Hello"))),
				},
				result: codes.OK,
			},
		},
		{
			name: "failed to download file: invalid userID",
			args: args{
				userID: "myUserID",
			},
			expects: expects{
				servAns: &model.File{
					FileName: "testfile.txt",
					FileSize: 100,
					Data:     *bufio.NewReader(bytes.NewReader([]byte("Hello"))),
				},
				wantErr: true,
				result:  codes.Unauthenticated,
			},
		},
		{
			name: "failed to download file: service error",
			args: args{
				userID: uuid.NewString(),
			},
			expects: expects{
				servErr: errors.New("file not found"),
				wantErr: true,
				result:  codes.InvalidArgument,
			},
		},
		{
			name: "failed to download file: file name is empty",
			args: args{
				userID: uuid.NewString(),
			},
			expects: expects{
				servAns: &model.File{
					FileName: "",
					FileSize: 100,
					Data:     *bufio.NewReader(bytes.NewReader([]byte("Hello"))),
				},
				wantErr: true,
				result:  codes.Internal,
			},
		},
		{
			name: "failed to download file: failed to send file",
			args: args{
				userID: uuid.NewString(),
			},
			expects: expects{
				servAns: &model.File{
					FileName: "testfile.txt",
					FileSize: 100,
					Data:     *bufio.NewReader(bytes.NewReader([]byte("Hello"))),
				},
				sendErr: errors.New("failed to send file"),
				wantErr: true,
				result:  codes.Internal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.WithValue(context.Background(), jwt.CtxUserIDKey, tt.args.userID)
			m := mock.NewMockServiceStorage(ctrl)
			m.EXPECT().DownloadFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.servAns, tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			stream := &grpcMock.ClientStreamMock[storage.DownloadRequest, storage.DownloadResponse]{
				SendFunc: func(s *storage.DownloadResponse) error {
					return tt.expects.sendErr
				},
				ContextFunc: func() context.Context {
					return ctx
				},
			}

			err := s.DownloadFile(&storage.DownloadRequest{FileId: tt.args.fileID}, stream)
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

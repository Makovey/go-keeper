package grpc

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/client/mock"
	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
)

func TestStorageClient_UploadFile(t *testing.T) {
	tmpFile, _ := os.CreateTemp("./", "file.txt")
	defer os.Remove(tmpFile.Name())

	type args struct {
		req  *storage.UploadRequest
		path string
	}

	type expects struct {
		sendErr   error
		clientErr error
		closeErr  error
		wantErr   bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "client successfully upload file",
			args: args{req: &storage.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{},
		},
		{
			name: "client failed to upload file: can't find file",
			args: args{req: &storage.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: "tmp.file.txt",
			},
			expects: expects{wantErr: true},
		},
		{
			name: "client successfully upload file",
			args: args{req: &storage.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{clientErr: errors.New("grpc client error"), wantErr: true},
		},
		{
			name: "client successfully upload file",
			args: args{req: &storage.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{sendErr: errors.New("can't send into stream"), wantErr: true},
		},
		{
			name: "client successfully upload file",
			args: args{req: &storage.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{closeErr: errors.New("stream closed with error"), wantErr: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := mock.ClientStreamMock{
				SendFunc: func(req *storage.UploadRequest) error {
					return tt.expects.sendErr
				},
				CloseAndRecvFunc: func() (*storage.UploadResponse, error) {
					return nil, tt.expects.closeErr
				},
			}

			m := mock.NewStorageClientMock(&stream, tt.expects.clientErr)
			client := NewStorageClient(dummy.NewDummyLogger(), m)

			err := client.UploadFile(context.Background(), tt.args.path)
			if tt.expects.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

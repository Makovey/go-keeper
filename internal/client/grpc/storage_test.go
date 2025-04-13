package grpc

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/client/mock"
	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	utilsMock "github.com/Makovey/go-keeper/internal/utils/mock"
)

func TestStorageClient_UploadFile(t *testing.T) {
	tmpFile, _ := os.CreateTemp("./", "file.txt")
	defer os.Remove(tmpFile.Name())

	type args struct {
		req  *pb.UploadRequest
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
			args: args{req: &pb.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{},
		},
		{
			name: "client failed to upload file: can't find file",
			args: args{req: &pb.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: "tmp.file.txt",
			},
			expects: expects{wantErr: true},
		},
		{
			name: "client failed to upload file: client error",
			args: args{req: &pb.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{clientErr: errors.New("grpc client error"), wantErr: true},
		},
		{
			name: "client failed to upload file: send stream error",
			args: args{req: &pb.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{sendErr: errors.New("can't send into stream"), wantErr: true},
		},
		{
			name: "client failed to upload file: can't close a stream",
			args: args{req: &pb.UploadRequest{
				FileName:  tmpFile.Name(),
				ChunkData: make([]byte, 0)},
				path: tmpFile.Name(),
			},
			expects: expects{closeErr: errors.New("stream closed with error"), wantErr: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stream := mock.ClientStreamMock[pb.UploadRequest, pb.UploadResponse]{
				SendFunc: func(req *pb.UploadRequest) error {
					return tt.expects.sendErr
				},
				CloseAndRecvFunc: func() (*pb.UploadResponse, error) {
					return nil, tt.expects.closeErr
				},
			}

			m := mock.NewStorageWithUploadStream(&stream, tt.expects.clientErr)
			dir := utilsMock.NewMockDirManager(ctrl)
			client := NewStorageClient(dummy.NewDummyLogger(), dir, m)

			err := client.UploadFile(context.Background(), tt.args.path)
			if tt.expects.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStorageClient_GetUsersFiles(t *testing.T) {
	type expects struct {
		clientResp *pb.GetUsersFileResponse
		clientErr  error
		wantErr    bool
	}

	tests := []struct {
		name    string
		expects expects
	}{
		{
			name: "client successfully get users files",
			expects: expects{
				clientResp: &pb.GetUsersFileResponse{
					Files: []*pb.UsersFile{
						{
							FileName: "file1.txt",
						},
					},
				},
			},
		},
		{
			name: "client failed to get users files: client error",
			expects: expects{
				clientErr: errors.New("grpc client error"),
				wantErr:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewStorageWitUsersFile(tt.expects.clientResp, tt.expects.clientErr)
			dir := utilsMock.NewMockDirManager(ctrl)
			client := NewStorageClient(dummy.NewDummyLogger(), dir, m)

			files, err := client.GetUsersFiles(context.Background())
			if tt.expects.wantErr {
				assert.Error(t, err)
				assert.Nil(t, files)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, files)
			}
		})
	}
}

func TestStorageClient_DownloadFile(t *testing.T) {
	file, _ := os.Create("temp")
	defer os.Remove("temp")

	type args struct {
		chunks []*pb.DownloadResponse
		fileID string
	}

	type expects struct {
		file          *os.File
		recvErr       error
		clientErr     error
		dirManagerErr error
		wantErr       bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "client successfully get users files",
			args: args{
				chunks: []*pb.DownloadResponse{
					{
						ChunkData: []byte("file1.txt"),
						FileName:  "file1.txt",
					},
					{
						ChunkData: []byte("file2.txt"),
					},
				},
			},
			expects: expects{
				file: file,
			},
		},
		{
			name: "client failed get users files: file name is empty",
			args: args{
				chunks: []*pb.DownloadResponse{
					{
						ChunkData: []byte("file1.txt"),
						FileName:  "",
					},
					{
						ChunkData: []byte("file2.txt"),
					},
				},
			},
			expects: expects{
				file:    file,
				wantErr: true,
			},
		},
		{
			name: "client failed get users files: stream receive error",
			expects: expects{
				file:    file,
				recvErr: errors.New("grpc client error"),
				wantErr: true,
			},
		},
		{
			name: "client failed get users files: can't create directory",
			args: args{
				chunks: []*pb.DownloadResponse{
					{
						ChunkData: []byte("file1.txt"),
						FileName:  "file1.txt",
					},
					{
						ChunkData: []byte("file2.txt"),
					},
				},
			},
			expects: expects{
				file:          file,
				dirManagerErr: errors.New("can't create directory"),
				wantErr:       true,
			},
		},
		{
			name: "client failed get users files: client error, can't get stream",
			expects: expects{
				clientErr: errors.New("grpc client error"),
				file:      file,
				wantErr:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stream := mock.ServerStreamClientMock[pb.DownloadResponse]{
				RecvFunc: func() func() (*pb.DownloadResponse, error) {
					count := 0
					return func() (*pb.DownloadResponse, error) {
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
			}

			m := mock.NewStorageWithDownloadStream(&stream, tt.expects.clientErr)

			dir := utilsMock.NewMockDirManager(ctrl)
			dir.EXPECT().CreateDir(gomock.Any(), gomock.Any()).Return(tt.expects.dirManagerErr).AnyTimes()
			dir.EXPECT().CreateFile(gomock.Any()).Return(tt.expects.file, tt.expects.dirManagerErr).AnyTimes()

			client := NewStorageClient(dummy.NewDummyLogger(), dir, m)
			err := client.DownloadFile(context.Background(), tt.args.fileID)
			if tt.expects.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStorageClient_DeleteFile(t *testing.T) {
	type args struct {
		fileID   string
		fileName string
	}
	type expects struct {
		clientErr error
		wantErr   bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name:    "client successfully deleted file files",
			expects: expects{},
		},
		{
			name: "client deleted file: client error",
			expects: expects{
				clientErr: errors.New("grpc client error"),
				wantErr:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewStorageWitEmptyMock(tt.expects.clientErr)
			dir := utilsMock.NewMockDirManager(ctrl)
			client := NewStorageClient(dummy.NewDummyLogger(), dir, m)

			err := client.DeleteFile(context.Background(), tt.args.fileID, tt.args.fileName)
			if tt.expects.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStorageClient_UploadPlainText(t *testing.T) {
	type args struct {
		content string
	}

	type expects struct {
		clientResp *pb.UploadPlainTextTypeResponse
		clientErr  error
		wantErr    bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "client successfully upload plain text",
			expects: expects{
				clientResp: &pb.UploadPlainTextTypeResponse{
					FileName: "file1.txt",
				},
			},
		},
		{
			name: "client upload plain text failed: client error",
			expects: expects{
				clientErr: errors.New("grpc client error"),
				wantErr:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewStorageWitUploadPlainText(
				tt.expects.clientResp,
				tt.expects.clientErr,
			)

			dir := utilsMock.NewMockDirManager(ctrl)
			client := NewStorageClient(dummy.NewDummyLogger(), dir, m)

			resp, err := client.UploadPlainText(context.Background(), tt.args.content)
			if tt.expects.wantErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp)
			}
		})
	}
}

package storage

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/config/stub"
	"github.com/Makovey/go-keeper/internal/repository/entity"
	"github.com/Makovey/go-keeper/internal/repository/mock"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
	utilsMock "github.com/Makovey/go-keeper/internal/utils/mock"
)

func Test_service_UploadFile(t *testing.T) {
	type args struct {
		file   *model.File
		userID string
	}

	type expects struct {
		storager  error
		repoError error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully save file with metadata",
			args: args{file: &model.File{
				FileName: "file.txt",
				FileSize: 50,
			}},
		},
		{
			name: "fail to upload file: can't create directory",
			args: args{file: &model.File{
				FileName: "file.txt",
				FileSize: 50,
			}},
			expects: expects{storager: errors.New("can't save file")},
			wantErr: true,
		},
		{
			name: "fail to upload file: repository error",
			args: args{file: &model.File{
				FileName: "file.txt",
				FileSize: 50,
			}},
			expects: expects{repoError: errors.New("can't save file metadata")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := mock.NewMockFileStorager(ctrl)
			storageMock.EXPECT().Save(tt.args.userID, tt.args.file.FileName, gomock.Any()).Return(tt.expects.storager).AnyTimes()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().SaveFileMetadata(gomock.Any(), gomock.Any()).Return(tt.expects.repoError).AnyTimes()

			cryptoMock := utilsMock.NewMockCrypto(ctrl)
			cryptoMock.EXPECT().EncryptReader(gomock.Any(), gomock.Any()).Return(bufio.NewReader(bytes.NewReader([]byte("encrypted"))), nil).AnyTimes()

			s := NewStorageService(repoMock, storageMock, cryptoMock, stub.NewStubConfig())
			got, err := s.UploadFile(context.Background(), *tt.args.file, tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				_, uuidErr := uuid.Parse(got)
				assert.NoError(t, err)
				assert.NoError(t, uuidErr)
			}
		})
	}
}

func Test_service_DownloadFile(t *testing.T) {
	type args struct {
		userID string
		fileID string
	}

	type expects struct {
		repoAns     *entity.File
		repoError   error
		storagerErr error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully download file",
			expects: expects{repoAns: &entity.File{
				FileName: "file.txt",
				FileSize: 50,
				Path:     "./file.txt",
			}},
		},
		{
			name:    "failed to download file: repository error",
			expects: expects{repoError: errors.New("unexpected error")},
			wantErr: true,
		},
		{
			name: "failed to download file: storage error",
			expects: expects{
				repoAns: &entity.File{
					FileName: "file.txt",
					FileSize: 50,
					Path:     "./file.txt",
				},
				storagerErr: errors.New("unexpected error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().GetFileMetadata(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.repoAns, tt.expects.repoError).AnyTimes()

			storageMock := mock.NewMockFileStorager(ctrl)
			storageMock.EXPECT().Get(gomock.Any()).Return([]byte("bytes from file"), tt.expects.storagerErr).AnyTimes()

			cryptoMock := utilsMock.NewMockCrypto(ctrl)
			cryptoMock.EXPECT().DecryptString(gomock.Any(), gomock.Any()).Return("bytes from file", nil).AnyTimes()

			s := NewStorageService(repoMock, storageMock, cryptoMock, stub.NewStubConfig())
			got, err := s.DownloadFile(context.Background(), tt.args.userID, tt.args.fileID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func Test_service_GetUsersFiles(t *testing.T) {
	type args struct {
		userID string
	}

	type expects struct {
		repoAns   []*entity.File
		repoError error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully get users files",
			expects: expects{
				repoAns: []*entity.File{
					{
						ID:       "1",
						FileName: "file.txt",
						FileSize: 50,
					},
					{
						ID:       "2",
						FileName: "fil2.txt",
						FileSize: 50,
					},
				},
			},
		},
		{
			name:    "failed to get users files: repository error",
			expects: expects{repoError: errors.New("unexpected error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().GetUsersFiles(gomock.Any(), tt.args.userID).Return(tt.expects.repoAns, tt.expects.repoError).AnyTimes()

			storageMock := mock.NewMockFileStorager(ctrl)

			cryptoMock := utilsMock.NewMockCrypto(ctrl)
			s := NewStorageService(repoMock, storageMock, cryptoMock, stub.NewStubConfig())
			got, err := s.GetUsersFiles(context.Background(), tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expects.repoAns), len(got))
			}
		})
	}
}

func Test_service_DeleteUsersFile(t *testing.T) {
	type args struct {
		userID   string
		fileID   string
		fileName string
	}

	type expects struct {
		repoError   error
		storagerErr error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name:    "successfully delete users files",
			expects: expects{},
		},
		{
			name:    "failed to delete users files: repository error",
			expects: expects{repoError: errors.New("unexpected error")},
			wantErr: true,
		},
		{
			name:    "failed to get users files: storage error",
			expects: expects{storagerErr: errors.New("unexpected error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().DeleteUsersFile(gomock.Any(), tt.args.userID, tt.args.fileID).Return(tt.expects.repoError).AnyTimes()

			storageMock := mock.NewMockFileStorager(ctrl)
			storageMock.EXPECT().Delete(fmt.Sprintf("%s/%s", tt.args.userID, tt.args.fileName)).Return(tt.expects.storagerErr).AnyTimes()

			cryptoMock := utilsMock.NewMockCrypto(ctrl)

			s := NewStorageService(repoMock, storageMock, cryptoMock, stub.NewStubConfig())
			err := s.DeleteUsersFile(context.Background(), tt.args.userID, tt.args.fileID, tt.args.fileName)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_service_UploadPlainText(t *testing.T) {
	type args struct {
		userID  string
		content string
	}

	type expects struct {
		cryptoRes   string
		cryptoErr   error
		storagerErr error
		repoError   error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name:    "successfully upload plain text",
			expects: expects{},
		},
		{
			name:    "failed to upload plain text: repository error",
			expects: expects{cryptoErr: errors.New("unexpected error")},
			wantErr: true,
		},
		{
			name:    "failed to upload plain text: repository error",
			expects: expects{repoError: errors.New("unexpected error")},
			wantErr: true,
		},
		{
			name:    "failed to upload plain text: storage error",
			expects: expects{storagerErr: errors.New("unexpected error")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().SaveFileMetadata(gomock.Any(), gomock.Any()).Return(tt.expects.repoError).AnyTimes()

			storageMock := mock.NewMockFileStorager(ctrl)
			storageMock.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.storagerErr).AnyTimes()

			cryptoMock := utilsMock.NewMockCrypto(ctrl)
			cryptoMock.EXPECT().EncryptString(gomock.Any(), gomock.Any()).Return(tt.expects.cryptoRes, tt.expects.cryptoErr).AnyTimes()

			s := NewStorageService(repoMock, storageMock, cryptoMock, stub.NewStubConfig())
			res, err := s.UploadPlainText(context.Background(), tt.args.userID, tt.args.content)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			}
		})
	}
}

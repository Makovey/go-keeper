package storage

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/repository/mock"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
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
				Data:     *bytes.NewReader([]byte("hello file")),
				FileName: "file.txt",
				FileSize: 50,
			}},
		},
		{
			name: "fail to upload file: can't create directory",
			args: args{file: &model.File{
				Data:     *bytes.NewReader([]byte("hello file")),
				FileName: "file.txt",
				FileSize: 50,
			}},
			expects: expects{storager: errors.New("can't save file")},
			wantErr: true,
		},
		{
			name: "fail to upload file: can't create directory",
			args: args{file: &model.File{
				Data:     *bytes.NewReader([]byte("hello file")),
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
			storageMock.EXPECT().Save(tt.args.userID, tt.args.file.FileName, tt.args.file.Data).Return(tt.expects.storager).AnyTimes()

			repoMock := mock.NewMockRepositoryStorage(ctrl)
			repoMock.EXPECT().SaveFileMetadata(gomock.Any(), gomock.Any()).Return(tt.expects.repoError).AnyTimes()

			s := NewStorageService(repoMock, storageMock, dummy.NewDummyLogger())
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

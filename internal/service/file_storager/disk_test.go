package file_storager

import (
	"bufio"
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/utils/mock"
)

func Test_diskStorager_Save(t *testing.T) {
	file, _ := os.Create("temp")
	defer os.Remove("temp")

	type args struct {
		path string
		name string
		buf  *bufio.Reader
	}

	type expects struct {
		file          *os.File
		createDirErr  error
		createFileErr error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully save file on disk",
			args: args{
				buf: bufio.NewReader(file),
			},
			expects: expects{
				file: file,
			},
		},
		{
			name: "failed to save file on disk: can't create dir",
			args: args{
				buf: bufio.NewReader(file),
			},
			expects: expects{
				file:         file,
				createDirErr: errors.New("can't create dir"),
			},
			wantErr: true,
		},
		{
			name: "failed to save file on disk: can't create file",
			args: args{
				buf: bufio.NewReader(file),
			},
			expects: expects{
				createFileErr: errors.New("can't create file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockDirManager(ctrl)
			m.EXPECT().CreateDir(gomock.Any(), gomock.Any()).Return(tt.expects.createDirErr).AnyTimes()
			m.EXPECT().CreateFile(gomock.Any()).Return(tt.expects.file, tt.expects.createFileErr).AnyTimes()

			s := NewDiskStorager(dummy.NewDummyLogger(), m)
			err := s.Save(tt.args.path, tt.args.name, tt.args.buf)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_diskStorager_Get(t *testing.T) {
	type args struct {
		path string
	}

	type expects struct {
		readFileRes []byte
		readFileErr error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully read file from a disk",
			expects: expects{
				readFileRes: []byte("test"),
			},
		},
		{
			name: "failed to read file from a disk: can't read file",
			expects: expects{
				readFileErr: errors.New("can't read file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockDirManager(ctrl)
			m.EXPECT().ReadFile(gomock.Any()).Return(tt.expects.readFileRes, tt.expects.readFileErr).AnyTimes()

			s := NewDiskStorager(dummy.NewDummyLogger(), m)
			got, err := s.Get(tt.args.path)
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

func Test_diskStorager_Delete(t *testing.T) {
	type args struct {
		path string
	}

	type expects struct {
		removeFileErr error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name:    "successfully remove file from a disk",
			expects: expects{},
		},
		{
			name: "failed to remove file from a disk: can't find file",
			expects: expects{
				removeFileErr: errors.New("can't remove file, file does not exist"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockDirManager(ctrl)
			m.EXPECT().RemoveFile(gomock.Any()).Return(tt.expects.removeFileErr).AnyTimes()

			s := NewDiskStorager(dummy.NewDummyLogger(), m)
			err := s.Delete(tt.args.path)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

package storage

import (
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
)

func TestServer_DeleteUsersFile(t *testing.T) {
	type args struct {
		userID string
	}

	type expects struct {
		servErr error
		wantErr bool
		result  codes.Code
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "successfully delete users file",
			args: args{userID: uuid.NewString()},
			expects: expects{
				result: codes.OK,
			},
		},
		{
			name: "failed to delete users file: invalid user id",
			args: args{userID: "user"},
			expects: expects{
				wantErr: true,
				result:  codes.Unauthenticated,
			},
		},
		{
			name: "failed to delete users file: invalid user id",
			args: args{userID: uuid.NewString()},
			expects: expects{
				servErr: errors.New("unexpected error"),
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
			m.EXPECT().DeleteUsersFile(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			_, err := s.DeleteUsersFile(ctx, &storage.DeleteUsersFileRequest{FileId: "1", FileName: "txt"})
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

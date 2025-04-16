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
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/service/mock"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func TestServer_GetUsersFile(t *testing.T) {
	type args struct {
		userID string
	}

	type expects struct {
		servErr error
		servAns []*model.ExtendedInfoFile
		wantErr bool
		result  codes.Code
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name: "successfully get users file",
			args: args{userID: uuid.NewString()},
			expects: expects{servAns: []*model.ExtendedInfoFile{
				{
					ID:       "1",
					FileName: "file1.txt",
					FileSize: "100 B",
				},
			},
				result: codes.OK,
			},
		},
		{
			name: "failed to get users file: invalid user id",
			args: args{userID: "user"},
			expects: expects{servAns: []*model.ExtendedInfoFile{
				{
					ID:       "1",
					FileName: "file1.txt",
					FileSize: "100 B",
				},
			},
				wantErr: true,
				result:  codes.Unauthenticated,
			},
		},
		{
			name: "failed to get users file: invalid user id",
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
			m.EXPECT().GetUsersFiles(gomock.Any(), gomock.Any()).Return(tt.expects.servAns, tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			res, err := s.GetUsersFile(ctx, &emptypb.Empty{})
			if tt.expects.wantErr {
				assert.Equal(t, tt.expects.result, status.Code(err))
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expects.result, codes.OK)
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

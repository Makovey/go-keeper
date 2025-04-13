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

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/service/mock"
)

func TestServer_UploadPlainTextType(t *testing.T) {
	type args struct {
		req    *pb.UploadPlainTextTypeRequest
		userID string
	}

	type expects struct {
		servAns string
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
			name: "successfully upload plain text",
			args: args{
				userID: uuid.NewString(),
				req: &pb.UploadPlainTextTypeRequest{
					Content: "test",
				},
			},
			expects: expects{
				result: codes.OK,
			},
		},
		{
			name: "failed to upload plain text: invalid userID",
			args: args{
				userID: "MineUserID",
				req: &pb.UploadPlainTextTypeRequest{
					Content: "test",
				},
			},
			expects: expects{
				wantErr: true,
				result:  codes.Unauthenticated,
			},
		},
		{
			name: "failed to upload plain text: service error",
			args: args{
				userID: uuid.NewString(),
				req: &pb.UploadPlainTextTypeRequest{
					Content: "test",
				},
			},
			expects: expects{
				servErr: errors.New("service error"),
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
			m.EXPECT().UploadPlainText(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.expects.servAns, tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			res, err := s.UploadPlainTextType(ctx, tt.args.req)
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

package auth

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service"
	"github.com/Makovey/go-keeper/internal/service/mock"
	grpcMock "github.com/Makovey/go-keeper/internal/transport/grpc/mock"
)

func TestServer_LoginUser(t *testing.T) {
	type args struct {
		req *pb.LoginRequest
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
			name:    "successfully login user",
			args:    args{req: &pb.LoginRequest{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{result: codes.OK},
		},
		{
			name:    "failed to login user: invalid email",
			args:    args{req: &pb.LoginRequest{Email: "testemail.ru", Password: "ttTest"}},
			expects: expects{wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to login user: invalid password",
			args:    args{req: &pb.LoginRequest{Email: "test@test.ru", Password: "test"}},
			expects: expects{wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to login user: user doesn't exists",
			args:    args{req: &pb.LoginRequest{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{servErr: service.ErrUserNotFound, wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to login user: incorrect password",
			args:    args{req: &pb.LoginRequest{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{servErr: service.ErrIncorrectPassword, wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to login user: random error from service",
			args:    args{req: &pb.LoginRequest{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{servErr: service.ErrGeneratePassword, wantErr: true, result: codes.Internal},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockService(ctrl)
			m.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return("token", tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			ctx := grpc.NewContextWithServerTransportStream(context.Background(), &grpcMock.ServerTransportStreamMock{})
			_, err := s.LoginUser(ctx, tt.args.req)
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

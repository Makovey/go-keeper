package auth

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/service"
	"github.com/Makovey/go-keeper/internal/service/mock"
	grpcMock "github.com/Makovey/go-keeper/internal/transport/grpc/mock"
)

func TestServer_RegisterUser(t *testing.T) {
	type args struct {
		req *auth.User
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
			name:    "successfully register user",
			args:    args{req: &auth.User{Name: "ttName", Email: "tt@mail.ru", Password: "ttPassword"}},
			expects: expects{result: codes.OK},
		},
		{
			name:    "failed to register user: invalid name",
			args:    args{req: &auth.User{Name: "", Email: "tt@mail.ru", Password: "ttPassword"}},
			expects: expects{wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to register user: invalid email",
			args:    args{req: &auth.User{Name: "ttName", Email: "testEmail.ru", Password: "ttPassword"}},
			expects: expects{wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to register user: invalid password",
			args:    args{req: &auth.User{Name: "", Email: "tt@mail.ru", Password: "Pass"}},
			expects: expects{wantErr: true, result: codes.InvalidArgument},
		},
		{
			name:    "failed to register user: user already exists",
			args:    args{req: &auth.User{Name: "ttName", Email: "tt@mail.ru", Password: "ttPassword"}},
			expects: expects{servErr: service.ErrUserAlreadyExists, wantErr: true, result: codes.AlreadyExists},
		},
		{
			name:    "failed to register user: user already exists",
			args:    args{req: &auth.User{Name: "ttName", Email: "tt@mail.ru", Password: "ttPassword"}},
			expects: expects{servErr: service.ErrGeneratePassword, wantErr: true, result: codes.Internal},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockService(ctrl)
			m.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return("token", tt.expects.servErr).AnyTimes()

			s := &Server{
				log:     dummy.NewDummyLogger(),
				service: m,
			}

			ctx := grpc.NewContextWithServerTransportStream(context.Background(), &grpcMock.ServerTransportStreamMock{})
			_, err := s.RegisterUser(ctx, tt.args.req)
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

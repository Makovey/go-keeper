package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Makovey/go-keeper/internal/client/mock"
	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger/dummy"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func TestAuthClient_Login(t *testing.T) {
	type args struct {
		req *model.Login
	}

	type expects struct {
		clientAns *auth.AuthResponse
		clientErr error
		wantErr   bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name:    "client successfully login user",
			args:    args{req: &model.Login{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{clientAns: &auth.AuthResponse{Token: "testable-token"}},
		},
		{
			name:    "client fail to login user: grpc client returned error",
			args:    args{req: &model.Login{Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{clientErr: errors.New("grpc.client failed to login"), wantErr: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mock.NewAuthClientMock(
				tt.expects.clientAns,
				tt.expects.clientErr,
			)

			client := NewAuthClient(dummy.NewDummyLogger(), m)
			got, err := client.Login(context.Background(), tt.args.req)
			if tt.expects.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestAuthClient_Register(t *testing.T) {
	type args struct {
		req *model.User
	}

	type expects struct {
		clientErr error
		clientAns *auth.AuthResponse
		wantErr   bool
	}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{
			name:    "client successfully register user",
			args:    args{req: &model.User{Name: "TestableName", Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{clientAns: &auth.AuthResponse{Token: "testable-token"}},
		},
		{
			name:    "client fail to register user: grpc client returned error",
			args:    args{req: &model.User{Name: "TestableName", Email: "test@test.ru", Password: "ttTest"}},
			expects: expects{clientErr: errors.New("grpc.client failed to login"), wantErr: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mock.NewAuthClientMock(
				tt.expects.clientAns,
				tt.expects.clientErr,
			)

			client := NewAuthClient(dummy.NewDummyLogger(), m)
			got, err := client.Register(context.Background(), tt.args.req)
			if tt.expects.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

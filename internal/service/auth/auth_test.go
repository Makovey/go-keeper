package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/Makovey/go-keeper/internal/config/stub"
	"github.com/Makovey/go-keeper/internal/repository/entity"
	"github.com/Makovey/go-keeper/internal/repository/mock"
	serviceErr "github.com/Makovey/go-keeper/internal/service"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func Test_service_RegisterUser(t *testing.T) {
	type args struct {
		user *model.User
	}

	type expects struct {
		repoError error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name: "successfully generate new auth token when sign up",
			args: args{user: &model.User{Name: "ttName", Email: "tt@mail.ru", Password: "tt123"}},
		},
		{
			name:    "fail to generate new auth token when sign up: can't bcrypt hash password",
			args:    args{user: &model.User{Name: "ttName", Email: "tt@mail.ru", Password: strings.Repeat("1", 75)}},
			wantErr: true,
		},
		{
			name:    "fail to generate new auth token when sign up: user already exists",
			args:    args{user: &model.User{Name: "ttName", Email: "tt@mail.ru", Password: "tt123"}},
			expects: expects{repoError: serviceErr.ErrUserAlreadyExists},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepository(ctrl)
			repoMock.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(tt.expects.repoError).AnyTimes()

			cfg := stub.NewStubConfig()
			s := NewAuthService(repoMock, jwt.NewManager(cfg))
			got, err := s.RegisterUser(context.Background(), tt.args.user)
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

func Test_service_LoginUser(t *testing.T) {
	type args struct {
		login *model.Login
	}

	type expects struct {
		repoAnswer *entity.User
		repoError  error
	}

	tests := []struct {
		name    string
		args    args
		expects expects
		wantErr bool
	}{
		{
			name:    "successfully generate new auth token when log in",
			args:    args{login: &model.Login{Email: "tt@mail.ru", Password: "tt123"}},
			expects: expects{repoAnswer: &entity.User{ID: "1", Name: "ttName", Email: "tt@mail.ru", PasswordHash: "tt123"}},
		},
		{
			name:    "fail to generate new auth token when log in: email not found",
			args:    args{login: &model.Login{Email: "tt@mail.ru", Password: "tt123"}},
			expects: expects{repoAnswer: &entity.User{}, repoError: pgx.ErrNoRows},
			wantErr: true,
		},
		{
			name:    "fail to generate new auth token when log in: password doesn't match",
			args:    args{login: &model.Login{Email: "tt@mail.ru", Password: "tt123"}},
			expects: expects{repoAnswer: &entity.User{ID: "1", Name: "ttName", Email: "tt@mail.ru", PasswordHash: "random"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := mock.NewMockRepository(ctrl)
			repoMock.EXPECT().GetUserInfo(gomock.Any(), tt.args.login.Email).Return(tt.expects.repoAnswer, tt.expects.repoError).Times(1)

			pass, _ := bcrypt.GenerateFromPassword([]byte(tt.expects.repoAnswer.PasswordHash), bcrypt.DefaultCost)
			tt.expects.repoAnswer.PasswordHash = string(pass)

			cfg := stub.NewStubConfig()
			s := NewAuthService(repoMock, jwt.NewManager(cfg))
			got, err := s.LoginUser(context.Background(), tt.args.login)
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

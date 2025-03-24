package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

type AuthClient struct {
	cfg    config.Config
	log    logger.Logger
	client auth.AuthClient
}

func NewAuthClient(
	conn *grpc.ClientConn,
	log logger.Logger,
) *AuthClient {

	return &AuthClient{
		log:    log,
		client: auth.NewAuthClient(conn),
	}
}

func (a *AuthClient) Register(ctx context.Context, user *model.User) error {
	fn := "grpc.AuthClient.Register"

	_, err := a.client.RegisterUser(ctx, mapper.ToProtoFromUser(user))
	if err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	return nil
}

func (a *AuthClient) Login(ctx context.Context, user *model.Login) error {
	fn := "grpc.AuthClient.Login"

	_, err := a.client.LoginUser(ctx, mapper.FromProtoToLogin(user))
	if err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	return nil
}

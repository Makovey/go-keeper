package grpc

import (
	"context"
	"fmt"

	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/mapper"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

type AuthClient struct {
	log    logger.Logger
	client auth.AuthClient
}

func NewAuthClient(
	log logger.Logger,
	client auth.AuthClient,
) *AuthClient {

	return &AuthClient{
		log:    log,
		client: client,
	}
}

func (a *AuthClient) Register(ctx context.Context, user *model.User) (string, error) {
	fn := "grpc.AuthClient.Register"

	response, err := a.client.RegisterUser(ctx, mapper.ToProtoFromUser(user))
	if err != nil {
		return "", fmt.Errorf("[%s]: %v", fn, err)
	}

	return response.GetToken(), nil
}

func (a *AuthClient) Login(ctx context.Context, user *model.Login) (string, error) {
	fn := "grpc.AuthClient.Login"

	response, err := a.client.LoginUser(ctx, mapper.FromProtoToLogin(user))
	if err != nil {
		return "", fmt.Errorf("[%s]: %v", fn, err)
	}

	return response.GetToken(), nil
}

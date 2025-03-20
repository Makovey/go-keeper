package auth

import (
	"context"

	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

type Service interface {
	RegisterUser(ctx context.Context, user *model.User) (string, error)
	LoginUser(ctx context.Context, user *model.Login) (string, error)
}

type Server struct {
	auth.UnimplementedAuthServer

	log     logger.Logger
	service Service
}

func NewAuthServer(
	log logger.Logger,
	service Service,
) *Server {
	return &Server{
		log:     log,
		service: service,
	}
}

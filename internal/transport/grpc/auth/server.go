package auth

import (
	"context"

	pb "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

//go:generate mockgen -source=server.go -destination=../../../service/mock/auth_service_mock.go -package=mock
type Service interface {
	RegisterUser(ctx context.Context, user *model.User) (string, error)
	LoginUser(ctx context.Context, user *model.Login) (string, error)
}

type Server struct {
	pb.UnimplementedAuthServer

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

package auth

import (
	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc"
)

type Server struct {
	auth.UnimplementedAuthServer
	log     logger.Logger
	service grpc.Service
}

func NewAuthServer(
	log logger.Logger,
	service grpc.Service,
) *Server {
	return &Server{
		log:     log,
		service: service,
	}
}

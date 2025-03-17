package keeper

import (
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/transport/grpc"
)

type Repository interface {
}

type service struct {
	repo Repository
	cfg  config.Config
	log  logger.Logger
}

func NewService(
	repo Repository,
	cfg config.Config,
	log logger.Logger,
) grpc.Service {
	return &service{
		repo: repo,
		cfg:  cfg,
		log:  log,
	}
}

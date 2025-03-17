package app

import (
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
)

type App struct {
	cfg config.Config
	log logger.Logger
}

func NewApp(
	cfg config.Config,
	log logger.Logger,
) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() {}

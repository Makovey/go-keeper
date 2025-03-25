package main

import (
	syslog "log"

	"github.com/Makovey/go-keeper/internal/app"
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger/slog"
	"github.com/Makovey/go-keeper/internal/repository/postgres"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/service/keeper"
	"github.com/Makovey/go-keeper/internal/transport/grpc/auth"
)

func main() {
	log := slog.NewLogger()
	cfg := config.NewConfig(log)

	repo, err := postgres.NewPostgresRepo(cfg, log)
	if err != nil {
		syslog.Fatalf("[%s]: %s", "main", err.Error())
	}
	manager := jwt.NewManager(cfg)

	service := keeper.NewAuthService(repo, manager, cfg, log)
	authServer := auth.NewAuthServer(log, service)

	appl := app.NewApp(
		cfg,
		log,
		authServer,
	)

	appl.Run()

	if err = repo.Close(); err != nil {
		log.Errorf("closed repo with error: %s", err.Error())
	}
}

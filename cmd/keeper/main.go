package main

import (
	syslog "log"

	"github.com/Makovey/go-keeper/internal/app"
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger/slog"
	"github.com/Makovey/go-keeper/internal/repository/postgres"
	authService "github.com/Makovey/go-keeper/internal/service/auth"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/transport/grpc/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/storage"
)

func main() {
	log := slog.NewLogger()
	cfg := config.NewConfig(log)

	repo, err := postgres.NewPostgresRepo(cfg, log)
	if err != nil {
		syslog.Fatalf("[%s]: %s", "main", err.Error())
	}
	manager := jwt.NewManager(cfg)

	service := authService.NewAuthService(repo, manager, log)
	authServer := auth.NewAuthServer(log, service)
	storageServer := storage.NewStorageServer(log, service)

	appl := app.NewApp(
		cfg,
		log,
		authServer,
		storageServer,
	)

	appl.Run()

	if err = repo.Close(); err != nil {
		log.Errorf("closed repo with error: %s", err.Error())
	}
}

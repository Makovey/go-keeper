package main

import (
	syslog "log"

	"github.com/Makovey/go-keeper/internal/app"
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger/slog"
	"github.com/Makovey/go-keeper/internal/repository/postgres"
	authService "github.com/Makovey/go-keeper/internal/service/auth"
	"github.com/Makovey/go-keeper/internal/service/file_storager"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	storageService "github.com/Makovey/go-keeper/internal/service/storage"
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

	aService := authService.NewAuthService(repo, manager)
	sService := storageService.NewStorageService(
		repo,
		file_storager.NewDiskStorager(log),
	)

	authServer := auth.NewAuthServer(log, aService)
	storageServer := storage.NewStorageServer(log, sService)

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

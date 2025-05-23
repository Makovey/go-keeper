package main

import (
	"fmt"
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
	"github.com/Makovey/go-keeper/internal/utils"
)

var (
	buildVersion = "N/A" // версия приложения
	buildDate    = "N/A" // дата сборки
	buildCommit  = "N/A" // коммит сборки
)

func main() {
	log := slog.NewLogger()
	cfg := config.NewConfig(log)

	log.Info(fmt.Sprintf("build version: %s", buildVersion))
	log.Info(fmt.Sprintf("build date: %s", buildDate))
	log.Info(fmt.Sprintf("build commit: %s", buildCommit))

	repo, err := postgres.NewPostgresRepo(cfg)
	if err != nil {
		syslog.Fatalf("[%s]: %s", "main", err.Error())
	}
	manager := jwt.NewManager(cfg)

	aService := authService.NewAuthService(repo, manager)
	sService := storageService.NewStorageService(
		repo,
		file_storager.NewDiskStorage(log, utils.NewDirManager()),
		utils.NewCrypto(),
		cfg,
	)

	authServer := auth.NewAuthServer(log, aService)
	storageServer := storage.NewStorageServer(log, sService)

	appl := app.NewApp(
		cfg,
		log,
		authServer,
		storageServer,
	)

	err = appl.Run()
	if err != nil {
		log.Errorf("[%s]: %s", "main", err.Error())
	}

	if err = repo.Close(); err != nil {
		log.Errorf("closed repo with error: %s", err.Error())
	}
}

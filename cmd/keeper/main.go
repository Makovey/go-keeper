package main

import (
	"github.com/Makovey/go-keeper/internal/app"
	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger/slog"
)

func main() {
	cfg := config.NewConfig()
	log := slog.NewLogger()

	appl := app.NewApp(
		cfg,
		log,
	)

	appl.Run()
}

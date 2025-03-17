package slog

import (
	"fmt"
	"log/slog"
	"os"

	def "github.com/Makovey/go-keeper/internal/logger"
)

type logger struct {
	*slog.Logger
}

func (l *logger) Errorf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.Error(message)
}

func (l *logger) Infof(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.Info(message)
}

func NewLogger() def.Logger {
	return &logger{slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))}
}

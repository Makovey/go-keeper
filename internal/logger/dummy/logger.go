package dummy

import def "github.com/Makovey/go-keeper/internal/logger"

type logger struct{}

func NewDummyLogger() def.Logger {
	return &logger{}
}

func (l *logger) Debug(format string, args ...any) {}

func (l *logger) Info(format string, args ...any) {}

func (l *logger) Infof(format string, args ...any) {}

func (l *logger) Warn(format string, args ...any) {}

func (l *logger) Error(format string, args ...any) {}

func (l *logger) Errorf(format string, args ...any) {}

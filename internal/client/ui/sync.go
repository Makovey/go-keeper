package ui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Makovey/go-keeper/internal/gen/storage"
)

func (m *Model) syncCmd() tea.Cmd {
	return tea.Tick(30*time.Second, func(time.Time) tea.Msg { // TODO: to cfg
		return syncMsg{}
	})
}

type syncMsg struct{}

func (m *Model) loadDataCmd() tea.Cmd {
	return func() tea.Msg {
		if m.token == "" {
			return nil
		}

		data, err := m.storage.GetUsersFiles(m.setTokenToCtx(context.Background()))
		if err != nil {
			return errMsg{err}
		}
		return dataMsg{data}
	}
}

type dataMsg struct {
	data []*storage.UsersFile
}

type errMsg struct {
	err error
}

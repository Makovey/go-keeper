package ui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
)

func (m *Model) syncCmd() tea.Cmd {
	return tea.Tick(m.updateDuration, func(time.Time) tea.Msg {
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
	data []*pb.UsersFile
}

type errMsg struct {
	err error
}

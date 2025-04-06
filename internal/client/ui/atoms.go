package ui

import (
	"errors"
	"path/filepath"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultWidth = 40
	listHeight   = 12
)

func mainList(title string, items []list.Item) list.Model {
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.Styles.Title = lipgloss.NewStyle().Bold(true).Underline(true)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.DisableQuitKeybindings()

	return l
}

func nameInput() textinput.Model {
	n := textinput.New()
	n.Placeholder = "Name"
	n.CharLimit = 50
	n.Validate = validate

	return n
}

func emailInput() textinput.Model {
	e := textinput.New()
	e.Validate = validate
	e.CharLimit = 50
	e.Placeholder = "Email"

	return e
}

func passwordInput() textinput.Model {
	p := textinput.New()
	p.Placeholder = "Password"
	p.EchoMode = textinput.EchoPassword
	p.CharLimit = 50
	p.Validate = validate

	return p
}

func filePicker() filepicker.Model {
	p := filepicker.New()
	absPath, _ := filepath.Abs(".")
	p.CurrentDirectory = absPath
	p.ShowPermissions = false
	p.ShowHidden = true
	p.Height = 10

	return p
}

func tableContent() table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 40},
		{Title: "Name", Width: 20},
		{Title: "Size", Width: 8},
		{Title: "Created At", Width: 18},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func validate(input string) error {
	if utf8.RuneCountInString(input) < 5 {
		return errors.New("needs at least 5 characters")
	}
	return nil
}

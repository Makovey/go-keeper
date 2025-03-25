package ui

import (
	"errors"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultWidth = 40
	listHeight   = 12
)

func mainList() list.Model {
	items := []list.Item{
		item("Sign Up"),
		item("Sign In"),
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Welcome to keeper, choose option:"
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

func validate(input string) error {
	if utf8.RuneCountInString(input) < 5 {
		return errors.New("needs at least 5 characters")
	}
	return nil
}

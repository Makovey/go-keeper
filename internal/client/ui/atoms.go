package ui

import (
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

	return n
}

func emailInput() textinput.Model {
	e := textinput.New()
	e.Placeholder = "Email"
	e.TextStyle = focusedStyle
	e.PromptStyle = focusedStyle

	return e
}

func passwordInput() textinput.Model {
	p := textinput.New()
	p.Placeholder = "Password"
	p.EchoMode = textinput.EchoPassword

	return p
}

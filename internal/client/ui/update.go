package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	cancel = "ctrl+c"
	enter  = "enter"
	tab    = "tab"

	select1 = "Sign Up"
	select2 = "Sign In"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.startedPage.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case cancel:
			return m, tea.Quit
		case enter:
			switch m.step {
			case startedList:
				i, ok := m.startedPage.list.SelectedItem().(item)
				if !ok {
					return m, nil
				}

				choice := string(i)
				switch choice {
				case select1:
					m.step = signUp
					m.signUpPage.name.Focus()
				case select2:
					m.step = signIn
					m.signInPage.email.Focus()
				}
				return m, nil
			case signUp:
				// TODO: get data and make request
				return m, tea.Quit
			case signIn:
				// TODO: get data and make request
				return m, tea.Quit
			}

		case tab:
			switch m.step {
			case signUp:
				switch {
				case m.signUpPage.name.Focused():
					in := []*textinput.Model{&m.signUpPage.name, &m.signUpPage.password}
					makeActiveInput(&m.signUpPage.email, in)
				case m.signUpPage.email.Focused():
					in := []*textinput.Model{&m.signUpPage.name, &m.signUpPage.email}
					makeActiveInput(&m.signUpPage.password, in)
				case m.signUpPage.password.Focused():
					in := []*textinput.Model{&m.signUpPage.email, &m.signUpPage.password}
					makeActiveInput(&m.signUpPage.name, in)
				}
			case signIn:
				switch {
				case m.signInPage.email.Focused():
					in := []*textinput.Model{&m.signInPage.email}
					makeActiveInput(&m.signInPage.password, in)
				case m.signInPage.password.Focused():
					in := []*textinput.Model{&m.signInPage.password}
					makeActiveInput(&m.signInPage.email, in)
				}
			default:
				break
			}
		}
	}

	var cmd tea.Cmd
	switch m.step {
	case startedList:
		m.startedPage.list, cmd = m.startedPage.list.Update(msg)
	case signUp:
		switch {
		case m.signUpPage.name.Focused():
			m.signUpPage.name, cmd = m.signUpPage.name.Update(msg)
		case m.signUpPage.email.Focused():
			m.signUpPage.email, cmd = m.signUpPage.email.Update(msg)
		case m.signUpPage.password.Focused():
			m.signUpPage.password, cmd = m.signUpPage.password.Update(msg)
		}
	case signIn:
		switch {
		case m.signInPage.email.Focused():
			m.signInPage.email, cmd = m.signInPage.email.Update(msg)
		case m.signInPage.password.Focused():
			m.signInPage.password, cmd = m.signInPage.password.Update(msg)
		}
	}

	return m, cmd
}

func makeActiveInput(active *textinput.Model, inactive []*textinput.Model) {
	active.Focus()
	active.TextStyle = focusedStyle
	active.PromptStyle = focusedStyle

	for _, i := range inactive {
		i.Blur()
		i.TextStyle = noStyle
		i.PromptStyle = noStyle
	}
}

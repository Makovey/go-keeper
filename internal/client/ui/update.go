package ui

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc/metadata"
)

const (
	cancel = "ctrl+c"
	enter  = "enter"
	tab    = "tab"

	startedListSelect1 = "Sign Up"
	startedListSelect2 = "Sign In"

	mainMenuSelect1 = "Upload file"
	mainMenuSelect2 = "Download file"
	mainMenuSelect3 = "Delete file"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.startedPage.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		m.clientMessage = nil
		switch keypress := msg.String(); keypress {
		case cancel:
			switch m.step {
			case startedList:
				return m, tea.Quit
			case signUp:
				m.step = startedList
				return m, nil
			case signIn:
				m.step = startedList
				return m, nil
			case upload:
				m.step = mainMenu
				return m, nil
			default:
				return m, tea.Quit
			}
		case enter:
			switch m.step {
			case startedList:
				i, ok := m.startedPage.list.SelectedItem().(item)
				if !ok {
					return m, nil
				}

				choice := string(i)
				switch choice {
				case startedListSelect1:
					m.step = signUp
					makeActiveInput(&m.signUpPage.name, nil)
				case startedListSelect2:
					m.step = signIn
					makeActiveInput(&m.signInPage.email, nil)
				}
				return m, nil
			case signUp:
				if !m.isPageValid() {
					return m, nil
				}

				token, err := m.auth.Register(context.TODO(), m.GetRegisterUserData())
				if err != nil {
					m.clientMessage = err
					return m, nil
				}
				m.token = token
				m.step = mainMenu
				return m, nil
			case signIn:
				if !m.isPageValid() {
					return m, nil
				}

				token, err := m.auth.Login(context.TODO(), m.GetLoginData())
				if err != nil {
					m.clientMessage = err
					return m, nil
				}
				m.token = token
				m.step = mainMenu
				return m, nil
			case mainMenu:
				i, ok := m.mainMenuPage.list.SelectedItem().(item)
				if !ok {
					return m, nil
				}

				choice := string(i)
				switch choice {
				case mainMenuSelect1:
					m.step = upload
					return m, m.uploadPage.picker.Init()
				case mainMenuSelect2:
					//m.step = signIn
				case mainMenuSelect3:
					//m.step = upload
					//return m, m.uploadPage.picker.Init()
				}
				return m, nil
			case upload:
				if m.uploadPage.selectedFile != "" {
					err := m.storage.UploadFile(m.setTokenToCtx(context.TODO()), m.uploadPage.selectedFile)
					if err != nil {
						m.clientMessage = err
						return m, nil
					}
					m.clientMessage = errors.New("file uploaded successfully")
				}
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
			case upload:
				return m, nil
			default:
				break
			}
		}
		m.uploadPage.selectedFile = ""
	}

	return m, m.updateModelValue(msg)
}

func (m *Model) updateModelValue(msg tea.Msg) tea.Cmd {
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
	case mainMenu:
		m.mainMenuPage.list, cmd = m.mainMenuPage.list.Update(msg)
	case upload:
		m.uploadPage.picker, cmd = m.uploadPage.picker.Update(msg)

		if didSelect, path := m.uploadPage.picker.DidSelectFile(msg); didSelect {
			m.uploadPage.selectedFile = path
		}

		if didSelect, path := m.uploadPage.picker.DidSelectDisabledFile(msg); didSelect {
			m.clientMessage = errors.New(path + " is not valid.")
			m.uploadPage.selectedFile = ""
		}
	}

	return cmd
}

func (m *Model) setTokenToCtx(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{"jwt": m.token})
	return metadata.NewOutgoingContext(ctx, md)
}

func (m *Model) isPageValid() bool {
	switch m.step {
	case signUp:
		return isInputValid(m.signUpPage.name) &&
			isInputValid(m.signUpPage.email) &&
			isInputValid(m.signUpPage.password)
	case signIn:
		return isInputValid(m.signInPage.email) &&
			isInputValid(m.signInPage.password)
	default:
		return true
	}
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

func isInputValid(input textinput.Model) bool {
	return utf8.RuneCountInString(input.Value()) != 0
}

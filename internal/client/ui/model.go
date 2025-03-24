package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type step int

const (
	startedList step = iota
	signUp
	signIn
)

type starterPage struct {
	list list.Model
}

type signUpPage struct {
	name     textinput.Model
	email    textinput.Model
	password textinput.Model
}

type signInPage struct {
	email    textinput.Model
	password textinput.Model
}

type Model struct {
	step        step
	startedPage starterPage
	signUpPage  signUpPage
	signInPage  signInPage
}

func InitialModel() Model {
	l := mainList()
	name := nameInput()
	email := emailInput()
	password := passwordInput()

	return Model{
		step:        startedList,
		startedPage: starterPage{l},
		signUpPage:  signUpPage{name: name, email: email, password: password},
		signInPage:  signInPage{email: email, password: password},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

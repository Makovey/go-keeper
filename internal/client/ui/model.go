package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Makovey/go-keeper/internal/client/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
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
	client      *grpc.AuthClient
	clientErr   error
}

func InitialModel(client *grpc.AuthClient) *Model {
	l := mainList()
	name := nameInput()
	email := emailInput()
	password := passwordInput()

	return &Model{
		step:        startedList,
		startedPage: starterPage{l},
		signUpPage:  signUpPage{name: name, email: email, password: password},
		signInPage:  signInPage{email: email, password: password},
		client:      client,
	}
}

func (m *Model) GetRegisterUserData() *model.User {
	return &model.User{
		Name:     m.signUpPage.name.Value(),
		Email:    m.signUpPage.email.Value(),
		Password: m.signUpPage.password.Value(),
	}
}

func (m *Model) GetLoginData() *model.Login {
	return &model.Login{
		Email:    m.signInPage.email.Value(),
		Password: m.signInPage.password.Value(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

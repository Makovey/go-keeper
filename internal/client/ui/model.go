package ui

import (
	"github.com/charmbracelet/bubbles/filepicker"
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
	mainMenu
	upload
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

type mainMenuPage struct {
	list list.Model
}

type uploadPage struct {
	picker       filepicker.Model
	selectedFile string
}

type Model struct {
	step          step
	startedPage   starterPage
	signUpPage    signUpPage
	signInPage    signInPage
	mainMenuPage  mainMenuPage
	uploadPage    uploadPage
	auth          *grpc.AuthClient
	storage       *grpc.StorageClient
	token         string
	clientMessage error
}

func InitialModel(
	auth *grpc.AuthClient,
	storage *grpc.StorageClient,
) *Model {
	starterList := mainList(
		"Welcome to auth, choose option:",
		[]list.Item{
			item("Sign Up"),
			item("Sign In"),
		},
	)
	name := nameInput()
	email := emailInput()
	password := passwordInput()
	mainMenuList := mainList(
		"What need to do?",
		[]list.Item{
			item("Upload file"),
			item("Download file"),
			item("Delete file"),
		},
	)

	return &Model{
		step:         startedList,
		startedPage:  starterPage{starterList},
		signUpPage:   signUpPage{name: name, email: email, password: password},
		signInPage:   signInPage{email: email, password: password},
		mainMenuPage: mainMenuPage{list: mainMenuList},
		uploadPage:   uploadPage{picker: filePicker()},
		auth:         auth,
		storage:      storage,
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

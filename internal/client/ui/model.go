package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Makovey/go-keeper/internal/client/grpc"
	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

type step int

const (
	startedList step = iota
	signUp
	signIn
	mainMenu
	download
	deleted
	upload
	creditCardUpload
	uploadText
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

type downloadPage struct {
	contentTable table.Model
	usersFiles   []*pb.UsersFile
}

type deletePage struct {
	contentTable table.Model
	usersFiles   []*pb.UsersFile
}

type uploadPage struct {
	picker       filepicker.Model
	selectedFile string
}

type uploadCreditCardPage struct {
	form             []textinput.Model
	focused          int
	validationErrors map[int]string
}

type uploadTextPage struct {
	textArea textarea.Model
}

type Model struct {
	step                 step
	startedPage          starterPage
	signUpPage           signUpPage
	signInPage           signInPage
	mainMenuPage         mainMenuPage
	downloadPage         downloadPage
	deletePage           deletePage
	uploadPage           uploadPage
	uploadCreditCardPage uploadCreditCardPage
	uploadText           uploadTextPage
	auth                 *grpc.AuthClient
	storage              *grpc.StorageClient
	updateDuration       time.Duration
	token                string
	clientMessage        error
}

func InitialModel(
	auth *grpc.AuthClient,
	storage *grpc.StorageClient,
	updateDuration time.Duration,
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
			item("Credit Card Number"),
			item("Plain text"),
		},
	)
	tableContent := tableContent()

	return &Model{
		step:                 startedList,
		startedPage:          starterPage{starterList},
		signUpPage:           signUpPage{name: name, email: email, password: password},
		signInPage:           signInPage{email: email, password: password},
		mainMenuPage:         mainMenuPage{list: mainMenuList},
		downloadPage:         downloadPage{contentTable: tableContent},
		deletePage:           deletePage{contentTable: tableContent},
		uploadPage:           uploadPage{picker: filePicker()},
		uploadCreditCardPage: uploadCreditCardPage{form: creditCardModel(), validationErrors: map[int]string{}},
		uploadText:           uploadTextPage{textArea: textArea()},
		auth:                 auth,
		storage:              storage,
		updateDuration:       updateDuration,
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

func (m *Model) GetCreditCardData() string {
	var key string
	var content strings.Builder
	for i, v := range m.uploadCreditCardPage.form {
		switch i {
		case 0:
			key = "ccn"
		case 1:
			key = "exp"
		case 2:
			key = "cvv"
		}
		content.WriteString(fmt.Sprintf("%s %s ", key, v.Value()))
	}

	return content.String()
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.syncCmd(),
	)
}

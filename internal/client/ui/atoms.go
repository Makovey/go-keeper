package ui

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
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

const (
	ccn = iota
	exp
	cvv
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
		{Title: "Created At (UTC 0)", Width: 18},
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

func creditCardModel() []textinput.Model {
	var inputs = make([]textinput.Model, 3)
	inputs[ccn] = textinput.New()
	inputs[ccn].Placeholder = "4505 **** **** 1234"
	inputs[ccn].Focus()
	inputs[ccn].CharLimit = 20
	inputs[ccn].Width = 30
	inputs[ccn].Prompt = ""
	inputs[ccn].Validate = ccnValidator

	inputs[exp] = textinput.New()
	inputs[exp].Placeholder = "MM/YY "
	inputs[exp].CharLimit = 5
	inputs[exp].Width = 5
	inputs[exp].Prompt = ""
	inputs[exp].Validate = expValidator

	inputs[cvv] = textinput.New()
	inputs[cvv].Placeholder = "XXX"
	inputs[cvv].CharLimit = 3
	inputs[cvv].Width = 5
	inputs[cvv].Prompt = ""
	inputs[cvv].Validate = cvvValidator

	return inputs
}

func validate(input string) error {
	if utf8.RuneCountInString(input) < 5 {
		return errors.New("needs at least 5 characters")
	}
	return nil
}

func ccnValidator(s string) error {
	if len(s) > 16+3 {
		return fmt.Errorf("ccn is too long")
	}

	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("ccn is invalid")
	}

	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("ccn must separate groups with spaces")
	}

	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}

func expValidator(s string) error {
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("exp is invalid")
	}

	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("exp is invalid")
	}

	return nil
}

func cvvValidator(s string) error {
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("cvv is invalid, must be only digits")
	}

	if len(s) > 3 {
		return fmt.Errorf("cvv is too long")
	}

	return nil
}

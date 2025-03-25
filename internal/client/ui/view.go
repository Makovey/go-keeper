package ui

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle          = lipgloss.NewStyle().Padding(1, 2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle           = lipgloss.NewStyle()
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func (m *Model) View() string {
	var b strings.Builder

	switch m.step {
	case startedList:
		b.WriteString(m.startedPage.list.View())
	case signUp:
		b.WriteString(m.signUpPage.name.View() + "\n")
		b.WriteString(m.signUpPage.email.View() + "\n")
		b.WriteString(m.signUpPage.password.View() + "\n")
	case signIn:
		b.WriteString(m.signInPage.email.View() + "\n")
		b.WriteString(m.signInPage.password.View() + "\n")
	}
	b.WriteString(showErrorIfNeeded(m.clientErr))

	return docStyle.Render(b.String() + "\n")
}

func showErrorIfNeeded(err error) string {
	if err == nil {
		return ""
	}

	re := regexp.MustCompile(`desc\s*=\s*(.*)`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Render("\n>> " + matches[1])
	}

	return err.Error()
}

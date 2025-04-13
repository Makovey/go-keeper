package ui

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
)

var (
	docStyle          = lipgloss.NewStyle().Padding(1, 2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	errorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	successStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("30"))
	noStyle           = lipgloss.NewStyle()
	inputStyle        = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle     = lipgloss.NewStyle().Foreground(darkGray)
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
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
		b.WriteString(m.signUpPage.password.View() + "\n\n")
		b.WriteString(continueStyle.Render("-> ↹ to switch input") + "\n")
		b.WriteString(continueStyle.Render("-> ↵  when fields is filled") + "\n")
		b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n\n")
	case signIn:
		b.WriteString(m.signInPage.email.View() + "\n")
		b.WriteString(m.signInPage.password.View() + "\n\n")
		b.WriteString(continueStyle.Render("-> ↹ to switch input") + "\n")
		b.WriteString(continueStyle.Render("-> ↵  when fields is filled") + "\n")
		b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n\n")
	case mainMenu:
		b.WriteString(m.mainMenuPage.list.View())
	case download:
		updateTable(&m.downloadPage.contentTable, m.downloadPage.usersFiles)
		b.WriteString(m.downloadPage.contentTable.View() + "\n\n")
		if len(m.downloadPage.usersFiles) == 0 {
			b.WriteString(focusedStyle.Render("you haven't files yet") + "\n")
		} else {
			b.WriteString(continueStyle.Render("-> ↑ to navigate up") + "\n")
			b.WriteString(continueStyle.Render("-> ↓ to navigate down") + "\n")
			b.WriteString(continueStyle.Render("-> ↵ to download file") + "\n")
			b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n\n")
		}
	case deleted:
		updateTable(&m.deletePage.contentTable, m.deletePage.usersFiles)
		b.WriteString(m.deletePage.contentTable.View() + "\n\n")
		if len(m.deletePage.usersFiles) == 0 {
			b.WriteString(focusedStyle.Render("you haven't files yet") + "\n")
		} else {
			b.WriteString(continueStyle.Render("-> ↑ to navigate up") + "\n")
			b.WriteString(continueStyle.Render("-> ↓ to navigate down") + "\n")
			b.WriteString(continueStyle.Render("-> ↵ to delete file") + "\n")
			b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n\n")
		}
	case upload:
		if m.uploadPage.selectedFile == "" {
			b.WriteString("Pick a file: \n")
		} else {
			b.WriteString("Selected file: " +
				m.uploadPage.picker.Styles.Selected.Render(m.uploadPage.selectedFile) + "\n" +
				"<< - select this file one more time, to confirm uploading\n",
			)
		}
		b.WriteString("\n" + m.uploadPage.picker.View() + "\n\n")
		b.WriteString(continueStyle.Render("-> ← to UP into file hierarchy") + "\n")
		b.WriteString(continueStyle.Render("-> → to DOWN into file hierarchy") + "\n")
		b.WriteString(continueStyle.Render("-> ↵ to choose file") + "\n")
		b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n\n")
	case creditCardUpload:
		s := fmt.Sprintf(
			`
 %s
 %s

 %s  %s
 %s  %s
`,
			inputStyle.Width(30).Render("Card Number"),
			m.uploadCreditCardPage.form[ccn].View(),
			inputStyle.Width(6).Render("EXP"),
			inputStyle.Width(6).Render("CVV"),
			m.uploadCreditCardPage.form[exp].View(),
			m.uploadCreditCardPage.form[cvv].View(),
		) + "\n"

		b.WriteString(s)
		b.WriteString(continueStyle.Render("-> ↹ to switch input") + "\n")
		b.WriteString(continueStyle.Render("-> ↵  when field is filled") + "\n")
		b.WriteString(continueStyle.Render("-> ctrl+c to move back") + "\n")
	case uploadText:
		b.WriteString(m.uploadText.textArea.View() + "\n\n")
		b.WriteString(continueStyle.Render("-> ↵ to break the line") + "\n")
		b.WriteString(continueStyle.Render("-> ctrl+s to save text") + "\n\n")
	}

	b.WriteString(showMessageIfNeeded(m.clientMessage))

	return docStyle.Render(b.String() + "\n")
}

func updateTable(model *table.Model, data []*pb.UsersFile) {
	rows := make([]table.Row, 0, len(data))

	for _, file := range data {
		rows = append(rows, table.Row{
			file.FileId,
			file.FileName,
			file.FileSize,
			file.CreatedAt.AsTime().Format("2006-01-02 15:04"),
		})
	}

	model.SetRows(rows)
}

func showMessageIfNeeded(err error) string {
	if err == nil {
		return ""
	}

	re := regexp.MustCompile(`desc\s*=\s*(.*)`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 1 {
		return errorStyle.Render("\n>> " + matches[1] + "\n")
	}

	return successStyle.Render("\n>> " + err.Error() + "\n")
}

package pages

import (
	"fmt"
	gitlab "gitlab.com/gitlab-org/api/client-go"

	t "github.com/OliverKeefe/git-cleanse/core/types"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AuthPageModel struct {
	platform string
	cursor   int
	input    []textinput.Model
	err      error
}

type AuthResultMsg struct {
	Success  bool
	Platform string
	Error    error
}

const (
	email = iota
	token
)

func NewAuthPage(platform string) AuthPageModel {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[email] = textinput.New()
	inputs[email].Placeholder = "Your email"
	inputs[email].Focus()
	inputs[email].Prompt = "✉️  "

	inputs[token] = textinput.New()
	inputs[token].Placeholder = "Password or Access Token"
	inputs[token].EchoMode = textinput.EchoPassword
	inputs[token].Prompt = "🔑  "

	return AuthPageModel{
		platform: platform,
		input:    inputs,
		cursor:   0,
	}
}

func (authModel AuthPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return authModel, tea.Quit

		case "tab", "shift+tab", "enter", "down", "up":
			if msg.String() == "enter" && authModel.cursor == len(authModel.input)-1 {
				return authModel, func() tea.Msg {
					return t.AuthSubmittedMsg{
						Platform: authModel.platform,
						Email:    authModel.input[email].Value(),
						Token:    authModel.input[token].Value(),
					}
				}
			}

			if msg.String() == "up" || msg.String() == "shift+tab" {
				authModel.cursor--
			} else {
				authModel.cursor++
			}

			if authModel.cursor < 0 {
				authModel.cursor = 0
			} else if authModel.cursor >= len(authModel.input) {
				authModel.cursor = len(authModel.input) - 1
			}

			for i := 0; i < len(authModel.input); i++ {
				if i == authModel.cursor {
					authModel.input[i].Focus()
				} else {
					authModel.input[i].Blur()
				}
			}
		}
	}

	for i := range authModel.input {
		var cmd tea.Cmd
		authModel.input[i], cmd = authModel.input[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return authModel, tea.Batch(cmds...)
}

func (m AuthPageModel) View() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render(fmt.Sprintf("Authenticate with %s", m.platform))

	formBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(1, 2)

	exitHint := lipgloss.NewStyle().
		Faint(true).
		Italic(true).
		Render("Press Enter to Submit, CTRL+C to Quit.")

	content := fmt.Sprintf(
		"%s\n\n%s\n%s\n\n%s",
		title,
		m.input[email].View(),
		m.input[token].View(),
		exitHint,
	)

	return lipgloss.Place(
		80, 20,
		lipgloss.Center, lipgloss.Center,
		formBox.Render(content),
	)
}

func (authModel AuthPageModel) Init() tea.Cmd { return nil }

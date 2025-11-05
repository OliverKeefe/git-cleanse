package pages

import (
	"fmt"

	"github.com/OliverKeefe/git-cleanse/core/repos"
	t "github.com/OliverKeefe/git-cleanse/core/types"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type AuthMsg struct {
	Success      bool
	Platform     string
	GitlabClient *gitlab.Client
	User         *repos.GitlabUser
	Error        error
}

type AuthModel struct {
	platform string
	cursor   int
	inputs   []textinput.Model
	err      error
}

const (
	emailInput = iota
	tokenInput
)

func NewAuthModel(platform string) AuthModel {
	inputs := make([]textinput.Model, 2)

	inputs[emailInput] = textinput.New()
	inputs[emailInput].Placeholder = "Your Email"
	inputs[emailInput].Prompt = "✉️  "
	inputs[emailInput].Focus()

	inputs[tokenInput] = textinput.New()
	inputs[tokenInput].Placeholder = "Access Token"
	inputs[tokenInput].Prompt = "🔑  "
	inputs[tokenInput].EchoMode = textinput.EchoPassword

	return AuthModel{
		platform: platform,
		inputs:   inputs,
		cursor:   0,
	}
}

func (m AuthModel) Init() tea.Cmd { return nil }

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab", "shift+tab", "down", "up":
			// Move cursor
			if msg.String() == "up" || msg.String() == "shift+tab" {
				m.cursor--
			} else {
				m.cursor++
			}

			if m.cursor < 0 {
				m.cursor = 0
			} else if m.cursor >= len(m.inputs) {
				m.cursor = len(m.inputs) - 1
			}

			for i := 0; i < len(m.inputs); i++ {
				if i == m.cursor {
					m.inputs[i].Focus()
				} else {
					m.inputs[i].Blur()
				}
			}

		case "enter":
			if m.cursor == len(m.inputs)-1 {
				email := m.inputs[emailInput].Value()
				token := m.inputs[tokenInput].Value()

				// Attempt GitLab auth asynchronously
				return m, func() tea.Msg {
					client, user, err := repos.InitGitLab(email, token, nil)
					if err != nil {
						return AuthMsg{
							Success:  false,
							Platform: m.platform,
							Error:    err,
						}
					}
					return AuthMsg{
						Success:      true,
						Platform:     m.platform,
						GitlabClient: client,
						User:         user,
					}
				}
			}
		}

	case AuthMsg:
		if msg.Success && msg.Platform == "GitLab" {
			// ✅ Navigate to GitLab page
			return m, func() tea.Msg {
				return t.NavigateMsg{To: "GitLabPage"}
			}
		}
		m.err = msg.Error
		return m, nil
	}

	// Update text inputs
	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AuthModel) View() string {
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
		m.inputs[emailInput].View(),
		m.inputs[tokenInput].View(),
		exitHint,
	)

	if m.err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true)
		content += fmt.Sprintf("\n\n%s", errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
	}

	return lipgloss.Place(
		80, 20,
		lipgloss.Center, lipgloss.Center,
		formBox.Render(content),
	)
}

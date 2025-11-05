package pages

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LocalRepoModel struct {
	Title string
}

func NewLocalRepoModel() LocalRepoModel {
	return LocalRepoModel{Title: "Local Repo"}
}

func (m LocalRepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m LocalRepoModel) View() string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("5")).
		Padding(1, 2)

	return style.Render(fmt.Sprintf("%s", m.Title))
}

func (m LocalRepoModel) Init() tea.Cmd {
	return nil
}

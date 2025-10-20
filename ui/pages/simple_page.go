package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MODEL DATA

type SimplePage struct{ text string }

func NewSimplePage(text string) SimplePage {
	return SimplePage{text: text}
}

// VIEW

func (s SimplePage) Init() tea.Cmd { return nil }

func (s SimplePage) View() string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(1, 0)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	exitHintStyle := lipgloss.NewStyle().
		Faint(true).
		Italic(true).
		MarginTop(1)

	content := titleStyle.Render("This is some placeholder content")
	box := borderStyle.Render(s.text)

	return fmt.Sprintf("%s%s\n%s", box, content, exitHintStyle.Render("Press Ctrl+C to exit."))
}

// UPDATE
func (s SimplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	return s, nil
}

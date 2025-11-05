package pages

import (
	"fmt"
	"strings"

	t "github.com/OliverKeefe/git-cleanse/core/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartMenuModel struct {
	cursor   int
	choices  []string
	selected string
	width    int
	height   int
}

func NewMenuModel() StartMenuModel {
	return StartMenuModel{
		cursor: 0,
		choices: []string{
			"Local Repository",
			"Github",
			"GitLab",
		},
		selected: "Local Repository",
	}
}

func (menuModel StartMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		menuModel.width = msg.Width
		menuModel.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return menuModel, tea.Quit

		case "up", "k":
			if menuModel.cursor > 0 {
				menuModel.cursor--
			}

		case "down", "j":
			if menuModel.cursor < len(menuModel.choices)-1 {
				menuModel.cursor++
			}

		case "enter":
			menuModel.selected = menuModel.choices[menuModel.cursor]
			var target t.NavigateMsg
			switch menuModel.selected {
			case "GitLab":
				target = t.NavigateMsg{To: "GitLab"}

			case "Local Repository":
				target = t.NavigateMsg{To: "Local"}

			default:
				target = t.NavigateMsg{To: "StartMenu"}

			}
			return menuModel, func() tea.Msg {
				return target
			}

		}
	}
	return menuModel, nil
}

func (menuModel StartMenuModel) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	cursorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true)

	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("170")).
		Bold(true)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 4).
		Margin(1).
		Width(menuModel.width - 50).
		BorderForeground(lipgloss.Color("63"))

	var builder strings.Builder
	builder.WriteString(titleStyle.Render("Select an Option:"))
	builder.WriteString("\n\n")

	for i, choice := range menuModel.choices {
		cursor := " "
		lineStyle := lipgloss.NewStyle()

		if i == menuModel.cursor {
			cursor = cursorStyle.Render("❯")
			lineStyle = selectedStyle
		}

		builder.WriteString(fmt.Sprintf("%s %s\n", cursor, lineStyle.Render(choice)))
	}

	content := boxStyle.Render(builder.String())

	width := menuModel.width
	height := menuModel.height
	if width == 0 {
		width = 80
	}
	if height == 0 {
		height = 24
	}

	return lipgloss.Place(menuModel.width, menuModel.height, lipgloss.Center, lipgloss.Center, content)
}

func (menuModel StartMenuModel) Selected() string {
	return menuModel.selected
}

func (menuModel StartMenuModel) Init() tea.Cmd { return nil }

package main

import (
	"fmt"
	"os"

	"github.com/OliverKeefe/git-cleanse/core/ui/routes"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
	program := tea.NewProgram(routes.NewRootModel())
	if _, err := program.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

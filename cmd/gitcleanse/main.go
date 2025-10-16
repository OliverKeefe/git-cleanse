package main

import (
	"fmt"
	"os"

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
		fmt.Println("Error:", err)
	}
}

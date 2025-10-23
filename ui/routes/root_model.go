package routes

import tea "github.com/charmbracelet/bubbletea"

type RootModel struct {
	currentPage PageID
	pages       map[PageID]tea.Model
}
type PageID string

const (
	StartMenu PageID = "startMenu"
)

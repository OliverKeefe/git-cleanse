package routes

import (
	t "github.com/OliverKeefe/git-cleanse/core/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type RootModel struct {
	currentPage t.PageID
	pages       map[t.PageID]tea.Model
}

const (
	PageStartMenu t.PageID = "StartMenu"
	PageGitLab    t.PageID = "GitLab"
)

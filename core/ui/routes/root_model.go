package routes

import (
	t "github.com/OliverKeefe/git-cleanse/core/ui"
	"github.com/OliverKeefe/git-cleanse/core/ui/pages"
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

func NewRootModel() RootModel {
	return RootModel{
		currentPage: PageStartMenu,
		pages: map[t.PageID]tea.Model{
			PageStartMenu: pages.NewMenuModel(),
			PageGitLab:    pages.NewSimplePage("Gitlab Selected"),
		},
	}
}

func (rootModel RootModel) Init() tea.Cmd { return nil }

func (rootModel RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case t.NavigateMsg:
		rootModel.currentPage = msg.To
		return rootModel, nil
	}

	current := rootModel.pages[rootModel.currentPage]
	newModel, cmd := current.Update(msg)
	rootModel.pages[rootModel.currentPage] = newModel
	return rootModel, cmd
}

func (rootModel RootModel) View() string {
	return rootModel.pages[rootModel.currentPage].View()
}
